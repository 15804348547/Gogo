package main

import "fmt"

func main() {
	//数组初始化及其读取
	var arr01 [6] int
	arr01[0]=1
	arr01[1]=2
	arr01[2]=3
	arr01[3]=4
	arr01[4]=5
	arr01[5]=6

	//数组的其他书初始化方式
	var a [4] int//= var a = [4] int
	fmt.Println(a)
	var b = [4] int{100,200,300,400}
	fmt.Println(b)
	var c = [4] int{1:888,2:999}//指定下标存储
	fmt.Println(c)
	d:=[6] string{"aaa","bbb","ccc","ddd","eee","fff"}
	fmt.Println(d)

	//for循环遍历
	for i := 0; i < 6; i++ {
		fmt.Println(arr01[i])
	}
	fmt.Println(len(arr01),cap(arr01))
	//forr迭代遍历
	for index, value := range d {
		fmt.Println(index,value)
	}
	fmt.Println("------------------------------")
	//forr省略下标遍历
	for _, value := range d {
		fmt.Println(value)
	}
	fmt.Println("------------------------------")

	//go二维数组
	var arr02 = [3][4]int{{4,4,4,4},{5,5,5},{1,1}}
	fmt.Println(arr02)
	//二维数组遍历
	for i := 0; i < len(arr02); i++ {

		for q := 0; q < len(arr02[i]); q++ {
			fmt.Print(arr02[i][q])
		}
		fmt.Println("---")
	}
}
