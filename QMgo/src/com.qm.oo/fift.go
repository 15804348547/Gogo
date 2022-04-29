package main

import "fmt"

func main() {
	var arr [6] int
	arr[0]=200
	arr[1]=300
	for _, value := range arr {
		fmt.Println(value)
	}
}
