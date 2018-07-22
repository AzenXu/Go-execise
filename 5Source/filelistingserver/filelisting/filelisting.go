package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"github.com/gpmgo/gopm/modules/log"
	"fmt"
	"strings"
)

const rootUrl = "/list/"

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

type userErrorType string

func (e userErrorType) Message() string {
	return string(e)
}

func (e userErrorType) Error() string {
	return e.Message()
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, rootUrl) != 0 {
		return userErrorType("资源路径必须以" + rootUrl + "开头啦，傻蛋！")
	}

	path := "./5Source/filelistingserver/" + request.URL.Path[len(rootUrl):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(content)
	return nil
}

func ErrWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			recover := recover()
			if recover != nil {
				fmt.Println("遇到一个Panic：", recover)
				http.Error(writer, "这是一个神奇的错误", http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Warn("Error handling request: %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}