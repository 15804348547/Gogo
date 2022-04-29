package main

import "fmt"

func main() {
//go语言中字符串类型相当与byte型的数组  ”.....“和``来表示字符串

	var name1 = "李克花"
	fmt.Println(name1)
	var name2 = `李克花`
	fmt.Println(name2)

	fmt.Printf("name1's type:%T --- value:%s\n",name1,name1)
	fmt.Printf("name2's type:%T --- value:%s",name2,name2)
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
	//单引号赋值相当于ASCII中的byte值    双引号相当于字符串类型
	byt01:='A'
	str01:="A"
	var num int = 99
	fmt.Printf("num type:%T -------value:%d",num,num)
	fmt.Printf("byt01's type:%T --- value:%d\n",byt01,byt01)
	fmt.Printf("str01's type:%T --- value:%s",str01,str01)

}
