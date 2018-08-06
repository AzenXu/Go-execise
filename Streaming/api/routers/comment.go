package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func LoadComments(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid查comments，按时间排序
}

func PostComment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid增comment
}

func DeleteComment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid删comment
}
