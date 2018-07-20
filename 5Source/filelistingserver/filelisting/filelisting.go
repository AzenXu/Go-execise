package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"github.com/gpmgo/gopm/modules/log"
)

const rootUrl = "/list/"

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
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
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())
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