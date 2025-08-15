package main

import (
	"fmt"

	"github.com/Rajshah1103/go-by-example/concurrency"
)

func main() {
	fmt.Println("This is the start of the main function.")
	// concurrency.InitializeWorkerPool()
	// concurrency.RateLimiterFunc()
	concurrency.AtomicCounterImpl()
}
