package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func registerHTTPRouter() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", homeHandler)

	return r
}

func main() {
	r := registerHTTPRouter()
	http.ListenAndServe(":8080", r)
}
