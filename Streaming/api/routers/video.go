package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func LoadVideos(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过用户ID查videos
}

func LoadVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid查id
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid删除id
}
