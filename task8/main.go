package main

import (
	"fmt"
	"log"
)


func main() {
	var (
		num int64
		bit uint8
		pos uint8
	)

	if count, err := fmt.Scan(&num, &bit, &pos); count != 3 || err != nil {
		log.Fatalf("Some errors in input: count=%d and err=%s", count, err.Error())
	}
	

	if getBit(num, pos) != bit {
		num ^= (1 << pos)
	}

	fmt.Printf("%d; binary: %b", num, num)
}


func getBit(num int64, pos uint8) uint8 {
	return uint8(num&(1 << pos) >> pos)
}