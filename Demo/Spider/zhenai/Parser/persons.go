package parser

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"log"
	"regexp"
)

var citylistRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
var morePersonListRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)">([^<]+)</a>`)

func PickUpPersons(content []byte) (items []engine.Item) {
	personMatches := citylistRe.FindAllSubmatch(content, -1)

	for _, m := range personMatches {
		name := string(m[2])
		items = append(items, engine.Item{
			Name: name,
			Request: engine.Request{
				URL: string(m[1]),
				ParasFunc: func(bytes []byte) []engine.Item {
					return PickUpPersonInfo(bytes, name)
				},
			},
		})
	}

	personListMatches := morePersonListRe.FindAllSubmatch(content, -1)

	for _, m := range personListMatches {
		name := string(m[2])
		items = append(items, engine.Item{
			Name: name,
			Request: engine.Request{
				URL:       string(m[1]),
				ParasFunc: PickUpPersons,
			},
		})
	}

	log.Println(items)

	return items
}
