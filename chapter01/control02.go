package main

/**switch*/
func main() {
	demo1()
	demo2()
}

func demo1()  {
	s := 100
	switch {
	case s > 0:
		println("正数")
	case s < 0:
		println("负数")
	default:
		println("0")
	}
}

func demo2()  {
	s := 2
	switch s {
	case 1:
		println("正数")
	case 2:
		println("负数")
	default:
		println("0")
	}
}
