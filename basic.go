package main

import (
	"fmt"
	"io/ioutil"
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

	//	条件语句
	bounded(438)
	fileRead()
}

/*
 * 条件语句
 */

func grade(score int) string {
	var g string
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "B"
	case score < 90:
		g = "A"
	case score <= 100:
		g = "SSR"
	}
	return g
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("不支持的操作类型")
	}
	return result
}

func fileRead() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(contents)
	} else {
		fmt.Println(err)
	}
}

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
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
