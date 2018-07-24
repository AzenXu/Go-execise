package parser

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"fmt"
	"regexp"
)

var citylistRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)

func PickUpPersons(content []byte) (items []engine.Item) {
	matches := citylistRe.FindAllSubmatch(content, -1)



	for _, m := range matches {
		name := string(m[2])
		items = append(items, engine.Item {
			Name: name,
			Request:engine.Request{
				URL: string(m[1]),
				ParasFunc: func(bytes []byte) []engine.Item {
					return PickUpPersonInfo(bytes, name)
				},
			},
		})

		fmt.Printf("Person: %s,URL: %s\n", m[2], m[1])
		fmt.Println()
	}
	fmt.Println(len(matches))

	return items
}
