package task1

import "fmt"

type Human struct {


}

func (h *Human) SayHello() {
	fmt.Println("Hello!")
}

type Action struct {
	Human
}


func test() {
	a := Action{}
	a.SayHello()
}