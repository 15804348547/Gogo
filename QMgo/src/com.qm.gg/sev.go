package main

import "fmt"

func main() {
	//go语言的类型转换,支持小数据类型转换成大数据类型，反之不可以.只有兼容时才可以转换，而常数一般系统都会自动转换
	var a int= 8
	var b int8= 10

	a = int(b)
	fmt.Println(a,b)

	c:=4.1333
	var d int = int(c)
	fmt.Println(c,d)




}
