package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"fmt"
)

func Fetch(URL string) (result []byte, err error) {
	return loadHtml(URL)
}

func loadHtml(URL string) (result []byte, err error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn("被禁止了！wrong code is: %d", resp.StatusCode)
		return nil, fmt.Errorf("被禁止了！wrong code is: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)

	log.Warn(URL)
	utf8Result := transform.NewReader(resp.Body, encodingJudgment(bodyReader).NewDecoder())
	result, err =  ioutil.ReadAll(utf8Result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// 判断传入Reader的编码格式
func encodingJudgment(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)

	if err != nil {
		log.Error(err.Error())
		return unicode.UTF8
	}

	e, name, _ := charset.DetermineEncoding(bytes, "")

	fmt.Println("我们猜测，该网站的编码格式为：", name)
	fmt.Println()
	return e
}
