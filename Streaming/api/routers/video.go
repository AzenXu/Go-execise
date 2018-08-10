package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"daker.wang/Azen/Go-execise/dbops"
	"daker.wang/Azen/Go-execise/Streaming/api/response"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"github.com/gpmgo/gopm/modules/log"
	"encoding/json"
	"daker.wang/Azen/Go-execise/Streaming/api/utils"
)

func LoadVideos(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过用户ID查videos
	username := params.ByName("username")
	videos, e := dbops.ListAllVideos(username, 0, utils.CurrentTimestampSec()); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	videosInfo := &defs.VideosInfo{Videos:videos}
	bytes, e := json.Marshal(videosInfo); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	response.SendNormalResponse(w, string(bytes), http.StatusOK)
}

func LoadVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid查id
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid删除id
}
