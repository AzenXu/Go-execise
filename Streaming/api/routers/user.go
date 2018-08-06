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

	//TODO:登录@Azen
	log.Warn("注册成功，username: %v, pwd: %v", user.Username, user.Pwd)
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params){

}

func LoadUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params){}

func Destory(w http.ResponseWriter, r *http.Request, params httprouter.Params){}
