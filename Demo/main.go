package main

import (
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"io"
	"bufio"
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	//loadHtml()

	//regexFullTextTest()
	//regexUseSimble()
	regexPickUp()
}

func regexPickUp() {
	const text = "My email is azen@daker.wang"

	regex := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+)`)
	result := regex.FindAllStringSubmatch(text, -1)

	fmt.Println(result)
}

func regexUseSimble() {
	const text = "My email is azen@daker.wang"

	regex := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	result := regex.FindAllString(text, -1)

	fmt.Println(result)
}

func regexFullTextTest() {
	const text = "My email is azen@daker.wang"

	re, err := regexp.Compile("azen@daker.wang")

	if err != nil {
		panic(err)
	}

	result := re.FindString(text)
	log.Warn(result)
}

func loadHtml() {
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

// 判断传入Reader的编码格式
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