package main

import (
	"fmt"
)

func main() {
	var v interface{}

	fmt.Scan(&v)

	switch v.(type) {
	case int:
		fmt.Print("int")
	case string:
		fmt.Print("string")
	case bool:
		fmt.Print("bool")
	case chan interface{}:
		fmt.Print("chan interface")
	default:
		fmt.Print("Unknown type")
	}
}