package engine

type Request struct {
	URL string
	ParasFunc func([]byte) []Item
}

type Item struct {
	Request Request
	Name interface{}
}
