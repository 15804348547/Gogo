package main

import "fmt"

func main() {
//goto关键字
	LOOP://声明goto--》的标签位置继续执行   可以用于中断和逻辑跳转处理
	var a = 10
	for a<20 {
		if a==15 {
			a+=1
			goto LOOP
			goto DONE
		}
		fmt.Println(a)
		a++
	}
	DONE:
	fmt.Println("all done......")
}
