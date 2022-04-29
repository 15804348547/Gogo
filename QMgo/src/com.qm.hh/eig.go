package main

import "fmt"

func main() {

	//go语言的运算符
	a := 3
	b := 10
	//算术运算符
	fmt.Println(b - a)
	fmt.Println(b + a)
	fmt.Println(b % a)
	fmt.Println(b * a)
	fmt.Println(b / a)
	fmt.Println("---------------------------------------------------------------")
	//自增自减运算符
	c := 100
	c++
	fmt.Println(c)
	d := 100
	d--
	fmt.Println(d)
	fmt.Println("----------------------------------------------------------------")
	//关系运算符  关系运算后返回结果一般为布尔型
	e := 2
	f := 3
	g := 6
	fmt.Println(e > f)
	fmt.Println(e < f)
	fmt.Println(g >= f)
	fmt.Println(g <= f)
	fmt.Println(e == f)
	fmt.Println(e != f)
	fmt.Println("---------------------------------------------------------------")
	//逻辑运算符  and/or/nor
	aa := true
	bb := false
	fmt.Println(aa && bb)
	fmt.Println(aa || bb)
	fmt.Println(!aa)
	fmt.Println(!bb)
	fmt.Println("---------------------------------------------------------------")
	//位运算符   主要为二进制的位操作逻辑
	aaa := 60
	bbb := 13
	fmt.Printf("aaa-dec:%d,aaa-bin:%b\n", aaa, aaa)
	fmt.Printf("bbb-dec:%d,bbb-bin:%b\n", bbb, bbb)
	fmt.Println("---------------------------------------------------------------")
		//按位与&
	fmt.Printf("aaa-dec:%d,aaa-bin:%b\n", aaa, aaa)
	fmt.Printf("&aaa-dec:%d,&aaa-bin:%b\n", aaa&bbb, aaa&bbb)
	fmt.Println("---------------------------------------------------------------")
		//按位或|
	fmt.Printf("aaa-dec:%d,aaa-bin:%b\n", aaa, aaa)
	fmt.Printf("&aaa-dec:%d,&aaa-bin:%b\n", aaa|bbb, aaa|bbb)
	fmt.Println("---------------------------------------------------------------")
		//异或操作^
	fmt.Printf("aaa-dec:%d,aaa-bin:%b\n", aaa, aaa)
	fmt.Printf("&aaa-dec:%d,&aaa-bin:%b\n", ^aaa, ^aaa)
	fmt.Println("---------------------------------------------------------------")
		//位清空&^
	fmt.Printf("aaa-dec:%d,aaa-bin:%b\n", aaa, aaa)
	fmt.Printf("&aaa-dec:%d,&aaa-bin:%b\n", aaa&^bbb, aaa&^bbb)
	fmt.Println("---------------------------------------------------------------")
	//位移操作符 a转换为二进制<<向左位移b位   a转换为二进制>>向右位移b位
	cccc:=8
	fmt.Println(cccc)
		//按位左移
	fmt.Println(cccc<<2)
		//按位右移
	fmt.Println(cccc>>2)
	fmt.Println("---------------------------------------------------------------")
	//赋值运算符  =、+=、-=、*=、/=、%=、<<=、>>=、&=、|=、^=.............
	ttt:=2
	rrr:=3
	rrr=ttt
	fmt.Println(rrr)
	ttt+=rrr
	fmt.Println(ttt)
	ttt-=rrr
	fmt.Println(ttt)
	ttt*=rrr
	fmt.Println(ttt)
	ttt%=rrr
	fmt.Println(ttt)
	fmt.Println("..........................................")

}
