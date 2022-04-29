package main

import (
	"fmt"
)

func main() {

	//标准输出
	fmt.Print("单行输出")
	fmt.Println("换行输出")
	fmt.Printf("格式化字符串%d",1000)
	fmt.Println("-----------------------------------------")
	//标准输入
	var name string
	var age int64
	fmt.Println("请输入您的姓名：")
	fmt.Scan(&name)
	fmt.Println("请输入您的年龄：")
	fmt.Scan(&age)
	fmt.Printf("欢迎：%s回来！年龄：%d",name,age)


}
