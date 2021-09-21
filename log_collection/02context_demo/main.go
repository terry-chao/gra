package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker() {
	defer wg.Done()
	for {
		fmt.Println("work....")
	}
}

func main() {
	wg.Add(1)
	go worker()

	wg.Wait()

	fmt.Println("over")
}
