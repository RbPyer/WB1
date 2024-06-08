package main

import (
	"fmt"
	"atomicgo.dev/constraints"
)



func main() {
	sl := []int{11, 5, 73, 59, 99, 8, 44, 85, 6, 10}
	fmt.Print(quickSort(sl))

}


func quickSort[T constraints.Orderable](s []T) []T{
	if len(s) <= 1 {
		return s
	}
	elem := s[0]

	left, mid, right := GetSlices(s, elem)

	result := make([]T, 0)
	result = append(result, quickSort(left)...)
	result = append(result, mid...)
	result = append(result, quickSort(right)...)


	return result

}




func GetSlices[T constraints.Orderable](s []T, elem T) ([]T, []T, []T) {
	ltSlice := make([]T, 0)
	gtSlice := make([]T, 0)
	eqSlice := make([]T, 0)

	for _, item := range s {
		if item > elem {
			gtSlice = append(gtSlice, item)
		} else if item < elem {
			ltSlice = append(ltSlice, item)
		} else {
			eqSlice = append(eqSlice, item)
		}
	}
	return ltSlice, eqSlice, gtSlice
}


