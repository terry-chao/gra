package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exit bool

func worker() {
	defer wg.Done()
	for {
		fmt.Println("work....")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go worker()

	time.Sleep(time.Second * 3)

	exit = true
	wg.Wait()

	fmt.Println("over")
}
