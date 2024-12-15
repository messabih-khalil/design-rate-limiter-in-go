package main

import (
	"fmt"
	"time"
)



type bucket struct{
	size int
	refilerSize int
	// this should be in minute
	refilerRoundTime int
}






func (bucket *bucket) refilerHandler(){
	// Refile the bucket each time provided in the bucket
	for range time.Tick(time.Duration(bucket.refilerRoundTime) * time.Minute) {
		bucket.size = 5
	}
}


func initRateLiniter() *bucket{
	
	bucket := &bucket{
			size : 5,
			refilerSize : 5,
			refilerRoundTime : 1,
	}
	
	// Refile The Bucket Each 1 minute
	go bucket.refilerHandler()
	
	// return the bucket 
	
	return bucket
	
}


func (bucket *bucket) rateLimiter(){
	
	

	notAllowed := bucket.size == 0 
	
	if notAllowed {
		fmt.Println("\n You hit the limit for today")	
		return
	}
	
	fmt.Println("\n You Have " , bucket.size , " left")	
	bucket.size -= 1
}

