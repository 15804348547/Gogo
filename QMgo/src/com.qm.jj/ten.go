package main

import "fmt"

func main() {

	//if判断和if嵌套
	a:=200
	if a>10 {
		fmt.Println("此数字大于10")
	}else {
		fmt.Println("此数字小于10")
	}

	if a>=10 {
		fmt.Println("此数字大于等于10")
		if a>=199 {
			fmt.Println("此数字大于等于199")
		}
	}

	if num:=6; num>4 {
		fmt.Println("if的特殊写法，可以在if语句头初始化变量，此变量作用域为if嵌套内")
	}

	//switch判断
	var season string = "春"
	switch season {
	case "春":
		fmt.Println("现在是春天")
		break//中断当前分支，不再执行下一分支
	case "夏":
		fmt.Println("现在是夏天")
		fallthrough//继续下一分支
	case "秋":
		fmt.Println("现在是秋天")
	case "冬":
		fmt.Println("现在是冬天")
	default:
		fmt.Println("输入季节错误........")

	}

	switch {
	case true:
		fmt.Println("执行true分支")
	case false:
		fmt.Println("执行false分支")
	default:
		fmt.Println("执行默认分支")
	}

	letter := 'A'
	switch letter {
	case 'A','E','I','O','U':
		fmt.Println("letter是元音字母")
	case 'P','t','s':
		fmt.Println("辅音分支")
	default:
		fmt.Println("error!")
	}
	switch lang:="golang";lang {
	case lang:
		fmt.Println("switch分支声明变量")
	}

	//for循环
	for true {
		fmt.Println("无限循环")
		break
	}
	for nums:=0;nums<100;nums++ {
		fmt.Println(nums)
	}
	for nums:=0;nums<100;nums++ {
		for i := 100; i >=0 ; i-- {
			fmt.Println("嵌套第",i)
		}
		fmt.Println(nums)
	}

}
