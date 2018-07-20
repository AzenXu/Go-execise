package main

import (
	"net/http"

	"daker.wang/Azen/Go-execise/5Source/filelistingserver/filelisting"
)

func main() {
	rootUrl := "/list/"
	http.HandleFunc(rootUrl, filelisting.ErrWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}