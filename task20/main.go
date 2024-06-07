package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("Some errors in input: %s", err.Error())
	}
	log.Print(ChangeString(s))
}


func ChangeString(s string) string{
	sl := strings.Fields(s)
	reverseSlice(sl)
	return strings.Join(sl, " ")
}


func reverseSlice(s []string) {
	length := len(s)

	for i := 0; i < length / 2; i++ {
		s[i], s[length - 1 - i] = s[length - 1 - i], s[i]
	}
}
