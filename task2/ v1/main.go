package main

import (
	"fmt"
	"sync"
)


func main() {
	var arr [5]int = [5]int{2,4,6,8,10}
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go pow(i, &arr, &wg)
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Print(arr[i])
	}
 }


func pow(i int, arr *[5]int, wg *sync.WaitGroup) {
	arr[i] = arr[i] * arr[i]
	wg.Done()
}