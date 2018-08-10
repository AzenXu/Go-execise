package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"encoding/json"
	"github.com/gpmgo/gopm/modules/log"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"daker.wang/Azen/Go-execise/Streaming/api/response"
	"daker.wang/Azen/Go-execise/dbops"
	"daker.wang/Azen/Go-execise/Streaming/api/session"
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

	if len(user.Pwd) <= 0 || len(user.Username) <= 0 {
		log.Error("用户名密码不完整")
		response.SendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	//  存数据库
	_, err = dbops.Regist(user.Username, user.Pwd)

	if err != nil {
		log.Error("存数据库出错 - %v", err)
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	ss := session.GenerateSession(user.Username)

	responseSessionOK(w, ss.SessionID)

	log.Warn("注册成功，username: %v, pwd: %v", user.Username, user.Pwd)
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  拿到用户名和密码
	body, e := ioutil.ReadAll(r.Body); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	var user *defs.UserCredential

	err := json.Unmarshal(body, &user); if err != nil {
		log.Error(err.Error())
		response.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	//  通过用户名获取数据库中的密码
	pwd, err := dbops.QueryPwd(user.Username); if err != nil {
		log.Error(err.Error())
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	//  密码比较
	if pwd != user.Pwd {
		response.SendErrorResponse(w, defs.ErrorNotAuthUser)
		log.Warn("密码错误 - %v", user)
		return
	}
	//  分配session并返回resp
	ss := session.GenerateSession(user.Username)

	responseSessionOK(w, ss.SessionID)
}

func LoadUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//  通过用户名拿到用户
	name := params.ByName("username")
	uid, e := dbops.QueryUserID(name); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	//  返回用户信息
	log.Warn("👻 拿到了uid %v", uid)
	responseUIDOK(w, uid)
}

func Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过sessionID删除session
}

func Destory(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//  Logout
	//  SessionID拿到用户ID
	//  删除用户
}

func responseSessionOK(w http.ResponseWriter, sid string) {
	s := defs.SessionResult{ SessionID: sid, OK: true }
	ss, err := json.Marshal(s)
	if err != nil {
		log.Error("解码出错 - %v", err)
		return
	}

	response.SendNormalResponse(w, string(ss), http.StatusOK)
}

func responseUIDOK(w http.ResponseWriter, uid string) {
	ur := defs.UserResult{ UserID: uid, OK: true }
	urj, err := json.Marshal(ur)
	if err != nil {
		log.Error("解码出错 - %v", err)
		return
	}

	response.SendNormalResponse(w, string(urj), http.StatusOK)
}