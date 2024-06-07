package main

import (
	"fmt"
	"math"
)


func main() {
	var (
		temperatures [8]float64 = [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
		k int
	)
	data := make(map[int][]float64)

	for i := range temperatures {
		k = getKey(temperatures[i])
		data[k] = append(data[k], temperatures[i])
	}

	for key, value := range data {
		fmt.Printf("%d: %v\n", key, value)
	}

}


func getKey(num float64) int { return int(num - math.Mod(num, 10.0)) }