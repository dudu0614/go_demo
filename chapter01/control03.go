package main

/**for*/
func main() {
	//forDemo1()
	//forDemo2()
	//forDemo3()
	//forDemo4()
	forDemo5()
}

func forDemo1()  {
	for i := 0;i<10;i++ {
		println(i)
	}
}

func forDemo2()  {
	for i:=4;i>0;i-- {
		println(i)
	}
}

func forDemo3()  {
	s :=0
	for s < 5 {
		println(s)
		s++
	}
}

func forDemo4()  {
	x:=4
	for {
		println(x)
		x--
		if x<0{
			break
		}
	}
}

func forDemo5()  {
	 x := []int{100,101,102}
	 for i,n :=range x{
	 	println(i,":",n)
	 }
}