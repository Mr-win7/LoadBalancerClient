package main

import (
	"errors"
	"net/http"
	"time"
)

var retryTimes = 3

func sendRequest(url string) error {
	currentRetryTimes := retryTimes
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Host", DomainName)
	for currentRetryTimes > 0 {
		resp, err := (&http.Client{}).Do(req)
		INFO.Println(resp)
		if err == nil && (resp.StatusCode >= 200 && resp.StatusCode <= 500) {
			return nil
		}
		WARNING.Println(err.Error())

		timeout := (retryTimes-currentRetryTimes)*2 + 1
		time.Sleep(time.Duration(timeout))
		currentRetryTimes--
	}
	return errors.New("max of retries")
}

// SendStep Send the request to assigned backend.
// start several goroutines to control number of connections.
func SendStep(backendChannel <-chan Route, requestChannel chan<- int) {
	for i := 0; i < 5; i++ {
		go func() {
			for route := range backendChannel {
				backend := route.Backend
				err := sendRequest("http://" + backend.IP)
				if err != nil {
					WARNING.Println(err.Error())
					backend.SetIsAlive(false)
					requestChannel <- route.ID
				} else {
					WaitGroup.Done()
				}
			}
		}()
	}
}
