package main

import (
	"sync"
)

// DomainName Domain name for backends.
var DomainName = "www.ebay.com"

// WaitGroup Wait for request finish.
var WaitGroup sync.WaitGroup

var requestNum = 100

func main() {
	InitLog()
	IPListFetch(DomainName)
	backendChannel := make(chan Route)
	requestChannel := make(chan int, requestNum)

	RouteStep(requestChannel, backendChannel)
	SendStep(backendChannel, requestChannel)
	for i := 0; i < requestNum; i++ {
		requestChannel <- i
		WaitGroup.Add(1)
	}
	WaitGroup.Wait()

}
