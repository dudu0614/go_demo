package main

import "fmt"

func main() {
	maptest()
}
func maptest()  {
	m:=make(map[string]int)
	m["a"]=1
	x,ok := m["b"]
	fmt.Println(x,ok)
	fmt.Println(m)
	delete(m,"a")
	m["c"]=4
	fmt.Println(m)
}
