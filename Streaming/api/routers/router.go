package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func Register() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", HomeHandler)

	r.POST("/user", Regist)
	r.POST("/user/:username", Login)
	r.GET("/user/:username", LoadUserInfo)
	r.DELETE("/user/:username", Destory)

	r.GET("/user/:username/videos", LoadVideos)
	r.POST("/user/:username/videos", UploadVideo)
	r.GET("/user/:username/videos/:vid-id", LoadVideo)
	r.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	r.GET("/video/:vid/comments", LoadComments)
	r.POST("/video/:vid/comments", PostComment)
	r.DELETE("/video/:vid/comment/cid", DeleteComment)

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "泥豪不豪~~~")
}