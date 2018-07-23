package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/pkg/errors"
)

type MineError struct {
	content string
}

func (err MineError) Error() string {
	return "这是自定义的error"
}

func main() {
	//	defer
	// tryDefer()
	// tryWriteFile()
	// tryDeferWithParam()
	// errorDealOld()
	// errorDealProtect()
	errorCreate()
}

func errorCreate() {
	err := errors.New("这是一个自己创建的error")
	if err != nil {
		fmt.Println(err)
	}
}

func errorDealProtect() {
	openAndWriteFileProtect("./file.txt")
}

func openAndWriteFileProtect(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println("啊哦报错了:", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	content := "又是一个测试用的字符串"
	fmt.Fprintln(writer, content)
}

func errorDealOld() {
	// writeFile("./file.txt")
	openAndWirteFile("./file.txt")
}

func openAndWirteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("啊哦报错了:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	content := "又是一个测试用的字符串"
	fmt.Fprintln(writer, content)
}

func tryDeferWithParam() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 3 {
			panic("print too many")
		}
	}
}

func tryWriteFile() {
	writeFile("./test.txt")
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	content := "我是一个萌萌哒字符串...ㄟ..." +
		"\n然后...然后...听说我要被写入到文件中了，还是有些害羞呢..."
	fmt.Fprintln(writer, content)
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occured")
	fmt.Println(4)
}
