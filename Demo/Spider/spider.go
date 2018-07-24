package main

import (
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"daker.wang/Azen/Go-execise/Demo/Spider/zhenai/Parser"
)

const testURL = "http://album.zhenai.com/u/1545542317"

const testSuccessfulURL = "http://album.zhenai.com/u/1486671081"

func main() {
	var simpleEngine = new(engine.SimpleEngine)
	simpleEngine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun",
		ParasFunc: parser.PickUpCitys,
	})

	//resp, _ := http.Get(testURL)
	//content, _ := httputil.DumpResponse(resp, true)
	//fmt.Println(string(content))

}

func loadHtml() (result []byte, err error) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	utf8Result := transform.NewReader(resp.Body, encodingJudgment(resp.Body).NewDecoder())

	result, err =  ioutil.ReadAll(utf8Result)
	if err != nil {
		return nil, err
	}

	return result, nil
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
