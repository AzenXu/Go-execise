package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func LoadComments(w http.ResponseWriter, r *http.Request, params httprouter.Params){}

func PostComment(w http.ResponseWriter, r *http.Request, params httprouter.Params){}

func DeleteComment(w http.ResponseWriter, r *http.Request, params httprouter.Params){}
