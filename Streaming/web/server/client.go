package server

import (
	"log"
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
	"encoding/json"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func requestAPI(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	u, _ := url.Parse(b.Url)
	u.Host = "localhost" + ":" + u.Port()
	newUrl := u.String()

	log.Println("👻 准备转发请求，newURL:", newUrl, "Ori URL:", r.URL.RequestURI())

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", newUrl, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, err := http.NewRequest("POST", newUrl, bytes.NewBuffer([]byte(b.ReqBody)))
		if err != nil {
			log.Println(err)
		}
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", newUrl, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
		return
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}