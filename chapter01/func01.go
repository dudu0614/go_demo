package main

import (
	"errors"
	"fmt"
)

func main() {
	a,b := 10,0
	c,err := div(a,b)
	fmt.Println(c,err)
}

func div(a,b int) (int,error)  {
	if b==0{
		return 0, errors.New("divsion by zero")
	}
	return a / b, nil
}
