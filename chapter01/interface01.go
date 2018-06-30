package main

import "fmt"

type user2 struct {
	name string
	age byte
}

type Printer interface {
	Print()
}

func (u user2) Print()  {
	fmt.Println(u)
}

func main() {
	var u user2
	u.name="dudu0614"
	u.age=27
	u.Print()
}
