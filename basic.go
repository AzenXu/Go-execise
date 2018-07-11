package main

import (
	"fmt"
)

var (
	s string
	a, b int
	// a, b := 666, 233 // 全局变量不能用:=定义
)
var ss = "总想搞个大新闻"

func main() {
	fmt.Println("Hello Azen٩(●˙▿˙●)۶…⋆ฺ")
	variableZeroValueDefine()
	varibaleInitial()
	varibaleTypeDeduction()
	varibaleShortcut()
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

