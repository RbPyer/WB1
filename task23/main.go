package main

import (
	"fmt"
	"log"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int

	if count, err := fmt.Scan(&i); count != 1 || err != nil || i > len(s) {
		log.Fatalf("Some errors with your input: count - %d\nerr - %s, index - %d; len - %d",
			count, err.Error(), i, len(s))
	}

	s = append(s[:i], s[i+1:]... )
	fmt.Print(s)
}