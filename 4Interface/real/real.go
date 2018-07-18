package real

import (
	"net/http"
	"net/http/httputil"
)

// Retriever 真的请求数据
type Retriever struct {
	Contents string
}

// Get 方法
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	return string(result)
}
