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
	"io/ioutil"
	"daker.wang/Azen/Go-execise/Streaming/api/middleware/auth"
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

func UploadVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !auth.ValidateSession(r) {
		log.Error("Unathorized user \n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Error(err.Error())
		response.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	vi, err := dbops.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	log.Warn("Author id : %d, name: %s \n", nvbody.AuthorId, nvbody.Name)

	if err != nil {
		log.Error("Error in AddNewVideo: %s", err)
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		response.SendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		response.SendNormalResponse(w, string(resp), 201)
	}
}

func LoadVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid查id
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//  通过vid删除id
}
