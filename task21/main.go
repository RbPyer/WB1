package main

import "fmt"


func main() {
	h := Human{Name: "Mike"}
	h.Work()

	hp := ProgrammingAdapter{Human: h}
	hp.WriteCode()
}


type Human struct {
	Name string
	Worker
}


type ProgrammingAdapter struct {
	Human
}

type Worker interface {
	Work()
}

type Programmer interface {
	WriteCode()
}


func (h *Human) Work() {
	fmt.Printf("%s is working at the factory...\n", h.Name)
}

func (ha *ProgrammingAdapter) WriteCode() {
	fmt.Printf("Now %s is writing code...\n", ha.Name)
}
