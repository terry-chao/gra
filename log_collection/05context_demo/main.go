package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// channel

func worker(ctx context.Context) {
	defer wg.Done()
LABEL:
	for {
		fmt.Println("work....")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 3)
	cancel()

	wg.Wait()

	fmt.Println("over")
}
