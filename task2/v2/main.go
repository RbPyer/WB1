package main

import (
	"fmt"
)


func main() {
	var arr [5]int = [5]int{2,4,6,8,10}
	c := make(chan int)
	defer close(c)

	for i := 0; i < 5; i++ {
		go pow(arr[i], c)
		fmt.Printf("%d ", <-c)
	}
 }


func pow(num int, c chan int) {
	c <- num * num
}