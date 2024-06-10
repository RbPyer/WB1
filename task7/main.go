package main

import (
	"fmt"
	"sync"
)


func main() {
	m := make(map[int]int, 100)

	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	wg.Add(100)

	for i := range 100 {
		go func(m map[int]int, num int, mu *sync.Mutex, wg *sync.WaitGroup) {
			mu.Lock()
			m[num] = num + 1
			wg.Done()
			mu.Unlock()			
		}(m, i, mu, wg)
	}

	wg.Wait()

	fmt.Println(m)
}