package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func main() {
	var n int 
	fmt.Scan(&n)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n) * time.Second)
	defer cancel()
	c := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg := new(sync.WaitGroup)
	wg.Add(2)


	go func(ctx context.Context, c chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
				case <- ctx.Done():
					fmt.Println("Context was cancelled while writing...")
					close(c)
					return
				case c <- r.Int():
					fmt.Println("New value in channel was received!")
			}
		}

	}(ctx, c, wg)

	
	go func(ctx context.Context, c chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				fmt.Println("Context was cancelled while reading...")
				return
			case x, ok := <-c:
				if !ok {
					fmt.Println("Channel was closed...")
					return
				}
				fmt.Println(x)
			}
		}
	}(ctx, c, wg)

	// gracefull shutdown
	<-ctx.Done()
	wg.Wait()
	fmt.Println("Go gracefully goes to sleep... z-z-z-z...")
}