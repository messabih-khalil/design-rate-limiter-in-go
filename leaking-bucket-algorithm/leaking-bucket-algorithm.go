
package main

import (
	"fmt"
	"time"
)

type queue struct {
	size                 int
	currentRequestNumber int
}

// Adds a request to the queue
func (q *queue) addRequestToQueue() {
	if q.currentRequestNumber < q.size {
		q.currentRequestNumber++
		fmt.Printf("Added request to queue. Current requests: %d\n", q.currentRequestNumber)
	} else {
		fmt.Println("Queue is full. Cannot add more requests.")
	}
}

// Processes a request from the queue
func (q *queue) produceRequest() {
	if q.currentRequestNumber == 0 {
		fmt.Println("No requests to process.")
		return
	}

	q.currentRequestNumber--
	fmt.Printf("Processed a request. Remaining requests: %d\n", q.currentRequestNumber)
}

func testLeakingBucket() {
	// Initialize the queue with a maximum size of 10
	q := &queue{
		size: 10,
	}

	// Goroutine to add a request every 2 seconds
	go func() {
		for range time.Tick(2 * time.Second) {
			q.addRequestToQueue()
		}
	}()

	// Goroutine to process a request every 5 seconds
	go func() {
		for range time.Tick(5 * time.Second) {
			q.produceRequest()
		}
	}()

	// Prevent the program from exiting
	select {}
}
