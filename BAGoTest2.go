package main

import (
	"fmt"
)

/*
多行注释
嘿嘿
*/

var a = "boai"
var b = "boaihome.com"
var c bool

var (
	d int
	e bool
)
var f, m int = 5, 8

const (
	length  int = 10
	width   int = 5
	x, y, z     = 111, false, "str"
)

var area int

const (
	// iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
	a1 = iota
	a2
	a3
	a4 = "haha"
	a5
	a6
	a7 = 188
	a8 = iota
	a9
)

func main() {
	fmt.Println("hello!")

	fmt.Println(a, b, c, d, e, f, m)
	fmt.Println("字符串好长啊 字符串好长啊 字符串好长啊 字符串好长啊 字符串好长啊 字符串好长啊")

	area = length * width
	fmt.Println("面积 =", area)
	println()
	println(x, y, z)

	fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8, a9)

}
