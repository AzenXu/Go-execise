package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"net/http/httputil"
	"net/url"
	"github.com/gpmgo/gopm/modules/log"
)

const filePath = "./videos/"
const max_file_length = 1024 * 1024 * 50

func Regist() (r *httprouter.Router) {
	r = httprouter.New()

	r.GET("/", homeHandler)

	r.POST("/", homeHandler)

	r.GET("/userhome", userHomeHandler)

	r.POST("/userhome", userHomeHandler)

	r.POST("/api", aipHandler)

	r.GET("/videos/:vid-id", proxyVideoHandler)

	r.POST("/upload/:vid-id", proxyUploadHandler)

	r.ServeFiles("/statics/*filepath", http.Dir("./template"))

	r.GET("/test", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		files, e := template.ParseFiles("./upload.html"); if e != nil {
			return
		}

		files.Execute(writer, nil)
	})

	return r
}

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	scookie, se := request.Cookie("session")
	ucookie, ce := request.Cookie("username")

	if se != nil || ce != nil || len(scookie.Value) <= 0 || len(ucookie.Value) <= 0 {
		//  弹到注册页
		temp, err := template.ParseFiles("./template/home.html"); if err != nil {
			return
		}
		temp.Execute(writer, &HomePage{Name:"小白痴"})
	} else {
		//  弹到个人主页
		http.Redirect(writer, request, "/userhome", http.StatusFound)
	}
}

func userHomeHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	scookie, se := request.Cookie("session")
	ucookie, ce := request.Cookie("username")

	if se != nil || ce != nil || len(scookie.Value) <= 0 || len(ucookie.Value) <= 0 {
		//  弹到注册页
		http.Redirect(writer, request, "/", http.StatusFound)
	} else {
		//  弹到个人主页
		temp, err := template.ParseFiles("./template/userhome.html"); if err != nil {
			return
		}
		temp.Execute(writer, &HomePage{Name:"血小板"})
	}
}

func aipHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jresp, e := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if e != nil {
		return
	}

	apiBody := &ApiBody{}
	je := json.Unmarshal(jresp, apiBody); if je != nil {
		return
	}
	requestAPI(apiBody, writer, request)
}

func proxyVideoHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	u, e := url.Parse("http://localhost:9001"); if e != nil {
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(writer, request)
}

func proxyUploadHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	u, e := url.Parse("http://localhost:9001"); if e != nil {
		log.Error("🙉 上传转发出错 %v", e)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(writer, request)
}