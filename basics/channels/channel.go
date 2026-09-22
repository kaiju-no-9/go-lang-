package main

import (
	"fmt"
	"math/rand"
	"time"
)
// example ofr sending on data to goroutines 
func proceses(numChan chan int) {
	for i := range numChan {
		fmt.Println("processing the value here", i)
	}
}

func main1() {
	numChan := make(chan int)

	go proceses(numChan)

	for i := 0; i < 10; i++ {
		numChan <- rand.Intn(10)
	}
	close(numChan)

	time.Sleep(time.Second)
}