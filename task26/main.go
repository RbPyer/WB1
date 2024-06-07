package main

import (
	"log"
	"fmt"
	"strings"
)


func main() {
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		log.Fatalf("Some errors in input: %s", err.Error())
	}

	if checkUnique(s) {
		fmt.Printf("All symbols in %s are unique.", s)
	} else {
		fmt.Printf("All symbols in %s are not unique.", s)
	}


}


func checkUnique(s string) bool {
	m := make(map[string]bool, len(s))
	var tmp string

	for i := 0; i < len(s); i++ {
		tmp = strings.ToLower(string(s[i]))
		if m[tmp] == true {
			return false
		}
		m[tmp] = true

	}
	return true
}