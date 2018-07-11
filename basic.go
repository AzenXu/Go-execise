package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
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

	//	循环语句
	accumulate()
	fmt.Println(convertToBin(4))
	// printFile("none")
	forever()

	//	函数
	result, _ := div(10, 7)
	fmt.Println(result)
	fmt.Println(
		apply(func(a, b int) int {
			return a + b
		}, 10, 20),
	)
}

/*
 * 函数
 */

func sum(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func apply(op func(a, b int) int, a, b int) (result int) {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func divNoRecommend(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

/*
 * 循环语句
 */
func forever() {
	for {
		fmt.Println("Bilibili")
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}

func convertToBin(number int) string {
	result := ""
	for ; number > 0; number /= 2 {
		lsb := number % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func accumulate() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	fmt.Println(sum)
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
