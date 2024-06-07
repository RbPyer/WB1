package main

import (
	"fmt"
)


func main() {
	sl1 := []string{"cat", "dog", "fish", "god"}
	sl2 := []string{"python", "dog", "kernel", "god"}

	fmt.Print(intersect(sl1, sl2))
}


func intersect[T comparable](sl1, sl2 []T) []T{
	buffer := make([]T, 0)

	for i := 0; i < len(sl1); i++ {
		if checkValue(sl2, sl1[i]) {
			buffer = append(buffer, sl1[i])
		}
	}

	return buffer
}


func checkValue[T comparable](sl []T, elem T) bool {
	for i := 0; i < len(sl); i++ {
		if sl[i] == elem { return true }
	}
	return false
}