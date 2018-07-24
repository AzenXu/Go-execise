package parser

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"regexp"
	"fmt"
)

var citiesre = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9A-Za-z]+)"[^>]*>([^<]+)</a>`)

func PickUpCitys(content []byte) (items []engine.Item) {
	matches := citiesre.FindAllSubmatch(content, -1)

	i := 0
	for _, m := range matches {
		items = append(items, engine.Item {
			Name: string(m[2]),
			Request:engine.Request{
				URL: string(m[1]),
				ParasFunc: PickUpPersons,
			},
		})

		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
		fmt.Println()
		if i > 1 {
			break
		}
		i++
	}
	fmt.Println(len(matches))

	return items
}
