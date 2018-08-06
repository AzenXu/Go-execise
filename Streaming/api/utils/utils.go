package utils

import (
	"time"
	"io"
	"crypto/rand"
	"fmt"
	"strconv"
)

//  当前时间 - 毫秒
func CurrentTimeMilli() int64 {
	return time.Now().UnixNano()/1000000
}

//  当前时间 - 秒
func CurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}

//  创建一个新的uuid
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
