package main

import (
	"fmt"
)

func main() {
	//变量显示声明声明方式1
	var age int
	age = 2333
	fmt.Printf("变量值：%d\n\r",age)
	//变量显示声明方式2
	var age2 int = 66666
	fmt.Println(age2)
	//变量推断声明
	var age3  = "6666.33"
	fmt.Printf("变量类型：%T,变量值：%s",age3,age3)
	fmt.Println()
	//变量的赋值声明
	age4:=666
	fmt.Printf("变量的数值是：%d\r",age4)
	//集体赋值
	var age5,age6,name1,name2 = 100.001,200,"李连杰","成龙"
	fmt.Println(age5,age6,name1,name2)
	var (
		name3 = "陈佩斯"
		name4 = "郭德纲"
		age7 = 77.7
	)
	fmt.Println(name3,name4,age7)

	var chars string = "ahsdkjahdksahkjsfjhjhf"
	fmt.Printf("content:%s\r",chars)
	fmt.Scan()

}
