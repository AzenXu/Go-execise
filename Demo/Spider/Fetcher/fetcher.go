package fetcher

import (
	"bufio"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(URL string) (result []byte, err error) {
	resp, err := LoadURL(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn("被禁止了！wrong code is: %d", resp.StatusCode)
		return nil, fmt.Errorf("被禁止了！wrong code is: %d", resp.StatusCode)
	}

	log.Warn(URL)

	bodyReader := bufio.NewReader(resp.Body)
	utf8Result := transform.NewReader(resp.Body, encodingJudgment(bodyReader).NewDecoder())
	result, err = ioutil.ReadAll(utf8Result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func LoadURL(URL string) (resp *http.Response, err error) {
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Error(err.Error())
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	return http.DefaultClient.Do(request)
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
