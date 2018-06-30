package main

type X int

func (x *X) incr()  {
	*x++
}

func main() {
	var x X
	x.incr()
	println(x)
}
