package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"encoding/json"
	"github.com/gpmgo/gopm/modules/log"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"daker.wang/Azen/Go-execise/Streaming/api/response"
	"daker.wang/Azen/Go-execise/Streaming/api/dbops"
)

func Regist(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ubody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("读取Body出错 - %s", err)
		response.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	var user defs.UserCredential

	err = json.Unmarshal(ubody, &user); if err != nil {
		log.Error("json序列化错误 - %s", err)
		response.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	//  存数据库
	_, err = dbops.Regist(user.Username, user.Pwd)

	if err != nil {
		log.Error("存数据库出错 - %v", err)
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	//TODO:分配session并返回resp@Azen
	log.Warn("注册成功，username: %v, pwd: %v", user.Username, user.Pwd)
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  拿到用户名和密码
	//  通过用户名获取数据库中的密码
	//  密码比较
	//  分配session并返回resp
}

func LoadUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//  通过用户名拿到用户
	//  返回用户信息
}

func Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过sessionID删除session
}

func Destory(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//  Logout
	//  SessionID拿到用户ID
	//  删除用户
}
