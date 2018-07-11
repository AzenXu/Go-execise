package main

import (
	"fmt"
	"math"
)

var (
	s     string
	a, bb int
	// a, b := 666, 233 // 全局变量不能用:=定义
)
var ss = "总想搞个大新闻"

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	fmt.Println("Hello Azen٩(●˙▿˙●)۶…⋆ฺ")
	//	常变量定义
	variableZeroValueDefine()
	varibaleInitial()
	varibaleTypeDeduction()
	varibaleShortcut()
	triangle()
	consts()
	enums()
}

/*
 * 常变量定义部分
 */

func enums() {
	const (
		azen = iota
		daker
		_
		buran
	)
	fmt.Println(azen, daker, buran)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func variableZeroValueDefine() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func varibaleInitial() {
	var a, b int = 1, 2
	var s string = "猜猜我是谁"
	fmt.Println(a, b, s)
}

func varibaleTypeDeduction() {
	var a, b, c, s = 1, 2, true, "ㄟ..."
	b = 666
	fmt.Println(a, b, c, s)
}

func varibaleShortcut() {
	a, b := 233, 666
	s := "老铁蛤蛤蛤"
	fmt.Println(a, b, s)
}
