package main

import (
  "sync"
)

var DomainName = "www.ebay.com"
var WaitGroup sync.WaitGroup

func main() {
  InitLog()
  IpListFetch(DomainName)
  backendChannel := make(chan *Backend, 100)
  requestChannel := make(chan int)

  RouteStep(requestChannel, backendChannel)
  SendStep(backendChannel, requestChannel)
  for i := 0; i < 100; i++ {
    requestChannel <- i
    WaitGroup.Add(1)
  }
  WaitGroup.Wait()

}
