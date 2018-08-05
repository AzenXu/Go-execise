package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("测试测试"))
}