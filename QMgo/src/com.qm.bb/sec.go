package main

import "fmt"

var name1 = "刘小傻"

func main() {
	var name = "李克花"
	fmt.Printf("被引用名称为：%s,被引用地址为：%p\n", name, &name)
	//内存对象的引用与存储值无关    改变赋值不会改变地址引用
	name = "曹德旺"
	fmt.Printf("被引用名称为：%s,被引用地址为：%p\n", name, &name)
	name = "李德全"
	fmt.Printf("被引用名称为：%s,被引用地址为：%p\n", name, &name)
	name = "叶赫那拉"
	fmt.Printf("被引用名称为：%s,被引用地址为：%p\n", name, &name)

	fmt.Println(name1)
	//空切片类型
	var s2 []int
	fmt.Println(s2)
	fmt.Println(s2 == nil)
	var ccc int64
	fmt.Println(ccc)
	var nnnn string
	fmt.Println(nnnn)
	var mmmm float64
	fmt.Println(mmmm)
	var iiii int8
	fmt.Println(iiii)
	type double float64
	type float float32
	var dddd float = 32.66666
	fmt.Println(dddd)
	var fffff float = 66.66
	fmt.Println(fffff)
}
