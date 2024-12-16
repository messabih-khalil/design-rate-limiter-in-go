package main

import "time"

func main() {
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
