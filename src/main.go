package main

import (
	"sync"
)

// DomainName Domain name for backends.
var DomainName = "www.ebay.com"
// WaitGroup Wait for request finish.
var WaitGroup sync.WaitGroup

func main() {
	InitLog()
	IPListFetch(DomainName)
	backendChannel := make(chan Route)
	requestChannel := make(chan int, 100)

	RouteStep(requestChannel, backendChannel)
	SendStep(backendChannel, requestChannel)
	for i := 0; i < 100; i++ {
		requestChannel <- i
		WaitGroup.Add(1)
	}
	WaitGroup.Wait()

}
