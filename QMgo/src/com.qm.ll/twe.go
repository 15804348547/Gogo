package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//go中随机数生成
	num01:=rand.Int()
	fmt.Println(num01)
	num02:=rand.Int31()
	fmt.Println(num02)
	num03:=rand.Float32()
	fmt.Println(num03)
	num04:=rand.Float64()
	fmt.Println(num04)

	//随机范围数
	num05:=rand.Int31n(100)
	fmt.Println(num05)
	rand.Seed(600)//设置种子数
	num06:=rand.Int63()
	fmt.Println(num06)

	//用时间戳作为随机数种子，相对意义上的随机数抽取方式
	tt:=time.Now()
	fmt.Println(tt)
	rand.Seed(tt.UnixNano())



}
