package main

import "fmt"

const PI = 3.1415926
const PATH = "https://www.baidu.com"

//成组常量声明  如果某一个变量没有初始化值   它的值将会等于上一个常量的值
const (
	MALE   = "男"
	FEMALE = "女"
	UNKNOWN
)

func main() {
	fmt.Println(MALE)
	fmt.Println(FEMALE)
	fmt.Println(UNKNOWN)
	fmt.Println(PI)
	const WNAME = "李克花"
	fmt.Printf("我的老婆叫：%s", WNAME)
	fmt.Println("")
	fmt.Printf("她的官网是%s", PATH)
	Square(2)

}

func Square(r float64) float64 {
	return r * r * PI / 2
}
