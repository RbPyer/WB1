package main

import (
	"fmt"
	"time"
)


func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Go to sleep for %d seconds...", n)
	sleep(time.Duration(n) * time.Second)

}


func sleep(d time.Duration) {
	t := time.After(d)
	<-t
}