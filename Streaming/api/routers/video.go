package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func LoadVideos(w http.ResponseWriter, r *http.Request, params httprouter.Params){}

func LoadVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params){}

func DeleteVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params){}
