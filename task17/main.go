package main

import (
	"fmt"
)


func main() {
	var sl []int = []int{1, 4, 6, 10, 13, 16, 20, 25, 47, 102, 208}

	var result int = binarySearch(sl, 102)
	if result == -1 {
		fmt.Print("No such element in slice.")
	} else {
		fmt.Print(result)
	}
}


func binarySearch(s []int, num int) int{
	var left, right, result int = 0, len(s) - 1, -1

	for ;left <= right && result == -1; {
		mid := (left + right) / 2
		if s[mid] > num {
			right = mid - 1
		} else if s[mid] < num {
			left = mid + 1
		} else {
			result = mid
		}
	}
	return result
}