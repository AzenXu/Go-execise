package routers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
	"io/ioutil"
	"daker.wang/Azen/Go-execise/Streaming/stream/response"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"github.com/gpmgo/gopm/modules/log"
	"os"
	"time"
	"html/template"
	"fmt"
)

const filePath = "./videos/"
const maxFileLength = 1024 * 1024 * 150

func Regist() (r *httprouter.Router) {
	r = httprouter.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		io.WriteString(writer, "哇又被你抓到啦~~羞羞羞👻")
	})

	r.GET("/videos/:vid-id", streamHandler)

	r.POST("/upload/:vid-id", uploadHandler)

	r.GET("/test", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		files, e := template.ParseFiles("./upload.html"); if e != nil {
			log.Error(e.Error())
			response.SendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
			return
		}

		files.Execute(writer, nil)
	})

	return r
}

func uploadHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//  检查大小
	request.Body = http.MaxBytesReader(writer, request.Body, maxFileLength)

	//  解析文件
	err := request.ParseMultipartForm(maxFileLength); if err != nil {
		response.SendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
	}

	//  拿存储路径
	fn := params.ByName("vid-id")
	uri := filePath + fn

	//  关键点：拿file
	/*
		1. 需要约定html侧，通过<form>标签传文件
		2. form标签的name属性值需为"file"
		3. 如上，即可通过request.FormFile(key)拿到

	*/
	formFile, _, e := request.FormFile("file"); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}

	bytes, e := ioutil.ReadAll(formFile); if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}

	//  写文件
	err = ioutil.WriteFile(uri, bytes, 0666); if err != nil {
		fmt.Println("我错了...", err)
		response.SendErrorResponse(writer, defs.ErrorInternalFaults)
		return
	}

	//  报告结果
	response.SendNormalResponse(writer, "upload successful", http.StatusCreated)
}

func streamHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//  找到文件
	fn := params.ByName("vid-id")
	uri := filePath + fn
	//  打开文件 - 不是读二进制哦ioutil.ReadAll()
	file, e := os.Open(uri)
	defer file.Close()

	if e != nil {
		log.Error(e.Error())
		response.SendErrorResponse(writer, defs.ErrorInternalFaults)
		return
	}

	//  写头
	writer.Header().Set("Content-Type", "video/mp4")

	//  写流
	http.ServeContent(writer, request, "", time.Now(), file)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//  找到文件
	//  删除文件
	//  报告结果
}