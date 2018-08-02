package parser

import (
	"io/ioutil"
	"testing"
)

func TestPickUpPersons(t *testing.T) {
	content, _ := ioutil.ReadFile("city_list.htm")

	items := PickUpPersons(content)

	const expectSize = 20

	if len(items) != expectSize {
		t.Errorf("数量不对！应该有：%d 个;实际有 %d 个 ", expectSize, len(items))
	}
}
