package main

import "fmt"

func main() {
	//切片 引用类型的数据  可以使数组初始化动态化
	//数组
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	//切片
	var slice1 []int
	fmt.Println(slice1)

	slice2 := []string{"4", "5", "6"}
	fmt.Println(slice2)
	slice2 = append(slice2, "666", "777", "888")
	for ee := 0; ee < len(slice2); ee++ {
		fmt.Print(slice2[ee], "-")
	}
	fmt.Println("")

	slice3 := make([]int, 0, 5)
	fmt.Println(slice3)

	//数组追加合并
	slice1=append(slice1,slice3...)
	fmt.Println(slice1)
}
