package main

import "fmt"

func main() {
	fmt.Printf("hello world!")

	// 定义变量，在Go中未使用的变量会报异常
	var a int = 10
	// 定义多个变量
	var (
		name string
		age  int
	)
	b, c := 20, 10
	// 30的值丢弃不要。_表示是丢弃的值。
	_, d := 30, b+c+a

	// 定义
}
