package session

import (
	"sync"
)

var sessions *sync.Map // sessions缓存

func init() {
	sessions = &sync.Map{}
}

func LoadSessionsFromDB() {

}
