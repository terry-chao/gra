package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// channel

func worker(ch <-chan struct{}) {
	defer wg.Done()
LABEL:
	for {
		select {
		case <-ch:
			break LABEL
		default:
			fmt.Println("work...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var exitChan = make(chan struct{})

	wg.Add(1)
	go worker(exitChan)
	exitChan <- struct{}{}
	time.Sleep(time.Second * 3)

	wg.Wait()

	fmt.Println("over")
}
