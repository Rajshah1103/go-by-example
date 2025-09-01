package main

import (
	"fmt"

	"github.com/Rajshah1103/go-by-example/channels"
)

func main() {
	fmt.Println("This is the start of the main function.")
	// concurrency.InitializeWorkerPool()
	// concurrency.RateLimiterFunc()
	// concurrency.AtomicCounterImpl()
	// concurrency.MutexCounterImpl()
	channels.ChannelsImpl()
}
