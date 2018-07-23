package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"io"
	"bufio"
)

func encodingJudgment(r io.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}

	e, name, _ := charset.DetermineEncoding(bytes, "")

	fmt.Println("我们猜测，该网站的编码格式为：", name)
	fmt.Println()
	return e
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	utf8Result := transform.NewReader(resp.Body, encodingJudgment(resp.Body).NewDecoder())

	result, err :=  ioutil.ReadAll(utf8Result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
}
