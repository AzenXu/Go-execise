package ossops

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"azen/config"
)

var EndPoint string
var AccessKeyId string
var AccessKeySecret string

func init() {
	EndPoint = config.GetOSSAddress()
	AccessKeyId = ""
	AccessKeySecret = ""
}

func UploadToOss(filename string, path string, bucketname string) (ok bool) {
	client, err := oss.New(EndPoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		// HandleError(err)
		return false
	}

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		// HandleError(err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500 * 1024, oss.Routines(3))
	if err != nil {
		// HandleError(err)
		return false
	}

	return true
}
