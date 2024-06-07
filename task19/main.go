package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"task19/funcs"
)


func main() {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("Some errors in input: %s", err.Error())
	}
	fmt.Print(funcs.ReverseString(s))
}



