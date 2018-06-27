package main // 每个源文件都属于包的一部分，在文件头用package声明所属的包名称

import "fmt" //用import导入标准库或者第三方包

func main() {
	fmt.Println("Hello World") //入口函数main没有参数，且必须放在main包中。
}
