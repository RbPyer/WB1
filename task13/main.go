package main

import (
	"log"
	"fmt"
)


func main() {
	var a, b int
	if count, err := fmt.Scan(&a, &b); count != 2 || err != nil {
		log.Fatalf("Some errors in input: count=%d and err=%s", count, err.Error())
	}

	// #1 sugar
	a, b = b, a

	fmt.Printf("#1: %d %d\n", a, b)

	// #2 math
	// a = a + b
	// b = a - b
	// a = a - b

	// fmt.Printf("#2: %d %d\n", a, b)
}