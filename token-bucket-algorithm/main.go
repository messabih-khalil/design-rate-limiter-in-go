package main

import (
	"fmt"
	"time"
)

func main() {

	usersBuckets := map[int]*bucket{
		1: initRateLiniter(),
		2: initRateLiniter(),
		3: initRateLiniter(),
	}

	for i := 0; i < 10; i++ {

	}

	ub := usersBuckets[1]

	go ub.sendRequest()

	// Keep the program running
	select {}

}

func (b *bucket) sendRequest() {
	for range time.Tick(1 * time.Second) {
		b.rateLimiter()
		fmt.Println("Request sent at", time.Now())
	}
}
