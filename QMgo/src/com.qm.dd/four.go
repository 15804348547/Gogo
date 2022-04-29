package main

import "fmt"

func main() {
	/*
		iota自增赋值
		在const组声明中   iota可以使声明的数据进行自增
	*/
	const (
		A = iota
		B = iota
		C = iota
		D = iota
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

	const (
		e = iota
		f
		g
	)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	const (
		SPRING = iota
		SUMMER
		AUTUMN
		WINTER
	)
	fmt.Println(SPRING,SUMMER,AUTUMN,WINTER)

	const (
		h = 666
		i = iota
		j
		k
	)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
