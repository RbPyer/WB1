package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {
	var (
		s atomic.Int32
		data [5]int32 = [5]int32{2, 4, 6, 8, 10}
	)
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go add(data[i], &s, wg)
	}

	wg.Wait()
	fmt.Println(s.Load())
}


func add(num int32, s *atomic.Int32, wg *sync.WaitGroup) {
	s.Add(num * num)
	wg.Done()
}