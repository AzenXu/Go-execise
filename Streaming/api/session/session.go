package session

import (
	"sync"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"daker.wang/Azen/Go-execise/Streaming/api/utils"
	"github.com/gpmgo/gopm/modules/log"
	"daker.wang/Azen/Go-execise/Streaming/api/dbops"
	"database/sql"
)

var sessions *sync.Map // sessions缓存

func init() {
	sessions = &sync.Map{}
}

func LoadSessionsFromDB() {

}

// sessions
func GenerateSession(username string) (session *defs.Session) {
	//  生成session
	var err error
	session = &defs.Session{}
	session.UserName = username
	session.SessionID, err = utils.GenerateUUID(); if err != nil {
		log.Error(err.Error())
	}
	session.TTL = int64(utils.CurrentTimestampSec() + defs.SessionTTLs)
	//  返回
	go updateSession(*session)

	return session
}

func updateSession(session defs.Session) {
	sessions.Store(session.SessionID, session.UserName)
	_, e := dbops.SelectSession(session.SessionID)
	if e == sql.ErrNoRows {
		dbops.RegistSession(session)
	} else if e == nil {
		dbops.UpdateSession(session)
	}
}
