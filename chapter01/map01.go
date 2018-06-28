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
	delete(m,"a")
	fmt.Println(m)
}
