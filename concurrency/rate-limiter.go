package concurrency

import (
	"fmt"
	"time"
)

func RateLimiterFunc() {
	reqestsChannel := make(chan int, 5) // Channel to simulate requests
	for i := 1; i <= 5; i++ {
		reqestsChannel <- i
	}
	close(reqestsChannel)

	limiter := time.Tick(200 * time.Millisecond) // Rate limit of 5 requests per second
	for req := range reqestsChannel {
		<-limiter // Wait for the rate limit
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
