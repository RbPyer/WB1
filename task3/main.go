package main

import (
	"fmt"
	"sync"
)


func main() {
	var (
		s int
		data [5]int = [5]int{2, 4, 6, 8, 10}
		mu sync.Mutex = sync.Mutex{}
		wg sync.WaitGroup = sync.WaitGroup{}
	)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go add(data[i], &s, &mu, &wg)
	}

	wg.Wait()
	fmt.Println(s)
}


func add(num int, s *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	*s += num * num
	mu.Unlock()
	wg.Done()
}