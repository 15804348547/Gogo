package main

import "fmt"

func main() {
	//数据类型位数一般默认和电脑所安装的操作系统一致     但是数据类型的具体位数可以自己按需指定
	var i64 int64 = 99999999999999
	fmt.Println(i64)
	var iu64 uint64 = 9999999999999999
	fmt.Println(iu64)
	var i32 int32 = 666666666
	fmt.Println(i32)
	var iu8 uint8 = 255
	fmt.Println(iu8)
	var i8 int8 = -128
	fmt.Println(i8)
	var flo32 float32 = 333.333333333333333
	fmt.Println(flo32)
	var flo64 = 333.33333333333333333333333333
	fmt.Println(flo64)
	type double float64
	var dd double = 66.666666666666666666666666
	fmt.Println(dd)
	fmt.Printf("dd--->type:%T", dd)
	var boo bool
	boo = true
	fmt.Println(boo)

	//unit8数据位数和byte一致   可以相互传递字符值
	var iiu8 uint8 = 127
	var b8 byte
	b8 = iiu8
	fmt.Println(iiu8, b8)

}
