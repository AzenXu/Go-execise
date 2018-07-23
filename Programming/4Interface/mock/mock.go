package mock

// Retriever 返回mock数据
type Retriever struct {
	Contents string
}

// Get 返回mock数据
func (r *Retriever) Get(url string) string {
	return r.Contents
}
