package main

import "fmt"

func main() {
	// cannot use "s" (type untyped string) as type int in argument to addition
	//fmt.Println(addition(1,"s")) // 无法编译通过
	fmt.Println(addition(1, 3))

	// 定义一个int变量
	var i int = 3
	println(i)
}

func addition(i int, i2 int) int {
	return i + i2
}
