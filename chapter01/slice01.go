package main

import "fmt"

func main() {
	sliceTest()
}
func sliceTest() {
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}
