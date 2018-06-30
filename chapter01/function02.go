package main

import "fmt"

type user1 struct {
	name string
	age byte
}

type manage1 struct {
	user user1
	title string
}

func (user user1) toString() string  {
	return fmt.Sprintf("%+v",user)
}

func main() {
	var u user1
	u.name="dudu0614"
	u.age=28
	fmt.Println(u.toString())
}
