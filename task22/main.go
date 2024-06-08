package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		a, b, tmp big.Int
		op string
	)
	fmt.Scan(&a, &b, &op)

	switch op {
	case "+":
		fmt.Print(tmp.Add(&a, &b))
	case "-":
		fmt.Print(tmp.Sub(&a, &b))
	case "*":
		fmt.Print(tmp.Mul(&a, &b))
	case "/":
		fmt.Print(tmp.Div(&a, &b))
	}

}

