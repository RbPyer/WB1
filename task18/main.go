package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



func main() {


	
	var c Counter = Counter{}
	wg := new(sync.WaitGroup)
	wg.Add(1000)
	for i := 1; i < 1001; i ++ {
		go func(c *Counter, num int32) {
			c.Value.Add(num)
			wg.Done()
		}(&c, int32(i))
	} 
	wg.Wait()
	fmt.Println(c.Value.Load())
}


type Counter struct {
	Value atomic.Int32
}