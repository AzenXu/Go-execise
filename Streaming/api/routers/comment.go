package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"daker.wang/Azen/Go-execise/Streaming/api/middleware/auth"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"encoding/json"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"daker.wang/Azen/Go-execise/dbops"
	"daker.wang/Azen/Go-execise/Streaming/api/response"
	"strconv"
)

func LoadComments(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid查comments，按时间排序
}

func PostComment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  鉴权
	if !auth.ValidateSession(r) {
		log.Error("发表评论鉴权失败\n")
		return
	}
	//  拿到vid和comment
	vid := params.ByName("vid")
	res, _ := ioutil.ReadAll(r.Body)

	var comment defs.CommentRequest
	json.Unmarshal(res, &comment)

	log.Warn(string(res))

	//  插入数据库
	//TODO uid通过session取，而非直接从请求里读。请求里的东东不安全
	aid, _ := strconv.Atoi(comment.AuthorId)
	err := dbops.AddNewComments(vid, aid, comment.Content); if err != nil {
		log.Error(err.Error())
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	response.SendNormalResponse(w, "OK", http.StatusOK)
}

func DeleteComment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid删comment
}
