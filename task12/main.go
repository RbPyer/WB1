package main


import (
	"fmt"
)

func main() {
	var sdata [5]string =  [5]string{"cat", "cat", "dog", "cat", "tree"}
	mdata := make(map[string]bool, 5)
	
	for _, elem := range sdata {
		mdata[elem] = true
	}

	for _, v := range mdata {
		fmt.Println(v)
	}

}