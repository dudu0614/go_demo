package main

func main() {
	a,b :=100,0
	fucnTest(a,b)
}

func fucnTest(a,b int)  {
	defer println("disposeing.......")
	println(a/b)
}
