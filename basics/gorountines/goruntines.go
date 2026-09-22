package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task", id, "is running")

}

// we use wg to wait for all the gorountines to finish
func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
   wg.Wait()

}
