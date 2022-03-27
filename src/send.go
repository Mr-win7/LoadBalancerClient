package main

import (
  "net/http"
  "errors"
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

    timeout := (retryTimes - currentRetryTimes) * 2 + 1
    time.Sleep(time.Duration(timeout))
    currentRetryTimes--
  }
  return errors.New("max of retries");
}

func SendStep(backendChannel <-chan *Backend, requestChannel chan<- int) {
  for i := 0; i < 5; i++ {
    go func() {
      for backend := range backendChannel {
        err := sendRequest("http://" + backend.Ip)
        if err != nil {
          WARNING.Println(err.Error())
          backend.SetIsAlive(false)
          requestChannel <- 101
        } else {
          WaitGroup.Done()
        }
      }
    }()
  }
}
