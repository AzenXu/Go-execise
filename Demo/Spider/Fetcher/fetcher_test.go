package fetcher

import (
	"testing"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
)

const testURL = "http://album.zhenai.com/u/1545542317"

const testSuccessfulURL = "http://album.zhenai.com/u/1486671081"

func TestGetTest(t *testing.T) {
	resp := GetTest(testSuccessfulURL)

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}

func TestFetch(t *testing.T) {

	content, err := Fetch(testURL)

	if err != nil {
		log.Error(err.Error())
	}

	conString := string(content)

	fmt.Println(conString)
}
