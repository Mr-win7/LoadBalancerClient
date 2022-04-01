package main

import (
	"errors"
	"sync/atomic"
)

// Route Bind the request and backend.
type Route struct {
	ID      int
	Backend *Backend
}

var currentIndex uint64 = 0

func getNextIndex() (uint64, error) {
	for {
		value := atomic.LoadUint64(&currentIndex)
		if len(backendList) == 0 {
			return 0, errors.New("empty backend list")
		}
		newValue := (value + 1) % uint64(len(backendList))
		if atomic.CompareAndSwapUint64(&currentIndex, value, newValue) {
			return newValue, nil
		}
	}
}

func getRouteRecord() (*Backend, error) {
	if len(backendList) == 0 {
		return nil, errors.New("empty backend list")
	}

	nextIndex, err := getNextIndex()
  if (err != nil) {
    return nil, err
  }
	for i := 0; i < len(backendList); i++ {
		finalIndex := (nextIndex + uint64(i)) % uint64(len(backendList))
		if backendList[finalIndex].GetIsAlive() {
			return &backendList[finalIndex], nil
		}
	}
	return nil, errors.New("no available backend")
}

// RouteStep Get route for request.
func RouteStep(requestChannel <-chan int, routeChannel chan<- Route) {
	go func() {
		for requestID := range requestChannel {
			INFO.Println(requestID)
			backend, err := getRouteRecord()
			if err != nil {
				ERROR.Println(err.Error())
				WaitGroup.Done()
			} else {
				routeChannel <- Route{requestID, backend}
			}
		}
	}()
}
