package main

import "fmt"

func main() {
	var m manage
	m.name = "dudu0614"
	m.age = 27
	m.title = "测试"

	fmt.Println(m)
}

type user struct {
	name string
	age byte
}

type manage struct {
	user
	title string
}


