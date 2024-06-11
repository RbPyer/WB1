package main

import (
	"fmt"
	"time"
	"context"
)


func main() {

	var num uint8
	fmt.Scan(&num)
	switch num {
	case 1:
		case1()
	case 2:
		case2()
	case 3:
		case3()
	}

}



func case1() {
	ch := make(chan int)

	go func(c <-chan int) {
		for {
			value, ok := <-c
			if !ok {
				fmt.Println("Channel was closed")
				return
			}
			fmt.Println(value)
		}
	}(ch)

	for i := range 100 {
		ch <- i
	}
	close(ch)
	time.Sleep(5 * time.Second)
}


func case2() {
	ch := make(chan bool)

	go func(ch chan bool) {
		for {
			select {
			case _, ok := <- ch:
				if !ok {
					return
				}
				fmt.Println("Stopping goroutine and closing channel...")
				close(ch)
			default:
				fmt.Println("Goroutine is working...")
				time.Sleep(1 * time.Millisecond)
			}
		}
	}(ch)
	time.Sleep(1 * time.Second)
	ch <- true

}


func case3() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping goroutine...")
				return
			default:
				fmt.Println("Goroutine is working...")
				time.Sleep(100 * time.Millisecond)
			}
			
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}