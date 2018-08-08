package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/ops"
	"github.com/gpmgo/gopm/modules/log"
	"daker.wang/Azen/Go-execise/Streaming/api/response"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
)

func Register() (r *httprouter.Router) {
	r = httprouter.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		io.WriteString(writer, "哈哈哈哈哈哈哈我是肺炎大魔王！！")
	})

	r.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return r
}

func vidDelRecHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	vid := params.ByName("vid-id")
	if len(vid) <= 0 {
		log.Error("id为空")
		response.SendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}

	//  添加「待删数据库表」中
	e := ops.AddRecord(vid); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(writer, defs.ErrorDBError)
		return
	}

	response.SendNormalResponse(writer, "已加入「垃圾箱」，将会定时清理", http.StatusOK)
}