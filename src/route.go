package main

import (
  "sync/atomic"
  "errors"
)

var currentIndex uint64 = 0;
func getNextIndex() uint64 {
  for {
    value := atomic.LoadUint64(&currentIndex)
    newValue := (value + 1) % uint64(len(backendList))
    if atomic.CompareAndSwapUint64(&currentIndex, value, newValue) {
      return newValue
    }
  }
}

func getRouteRecord() (*Backend, error) {
  nextIndex := getNextIndex()
  for i := 0; i < len(backendList); i++ {
    finalIndex := (nextIndex + uint64(i)) % uint64(len(backendList))
    if backendList[finalIndex].GetIsAlive() {
      return &backendList[finalIndex], nil
    }
  }
  return nil, errors.New("No available backend!")
}

func RouteStep(requestChannel <-chan int, backendChannel chan<- *Backend) {
  go func() {
    for request := range requestChannel {
      INFO.Println(request)
      backend, err := getRouteRecord()
      if err != nil {
        ERROR.Println(err.Error())
        WaitGroup.Done()
      } else {
        backendChannel <- backend
      }
    }
  }()
}
