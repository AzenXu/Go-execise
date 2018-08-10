package session

import (
	"sync"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"daker.wang/Azen/Go-execise/Streaming/api/utils"
	"github.com/gpmgo/gopm/modules/log"
	"daker.wang/Azen/Go-execise/dbops"
	"database/sql"
)

var sessions *sync.Map // sessions缓存

func init() {
	sessions = &sync.Map{} // key: sid value: session
}

func loadSessionsFromDB(sid string) *defs.Session {
	session, err := dbops.SelectSessionFromSid(sid); if err != nil {
		log.Error(err.Error())
		return nil
	}
	return session
}

func loadSessionFromCache(sid string) *defs.Session {
	session, ok := sessions.Load(sid); if !ok {
		return nil
	}

	sm, ok := session.(defs.Session); if ok {
		log.Warn("OK")
	}
	return &sm
}

func loadSession(sid string) (session *defs.Session) {
	session = loadSessionFromCache(sid)
	if session != nil {
		return session
	}

	session = loadSessionsFromDB(sid)
	if session != nil {
		sessions.Store(session.SessionID, session)
		return session
	}

	return nil
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

func IsSessionExpired(sid string) (ok bool) {
	session := loadSession(sid)
	if session == nil {
		return false
	}

	if session.TTL < int64(utils.CurrentTimestampSec()) {
		return false
	}

	return true
}

func updateSession(session defs.Session) {
	sessions.Store(session.SessionID, session)
	_, e := dbops.SelectSession(session.SessionID)
	if e == sql.ErrNoRows {
		dbops.RegistSession(session)
	} else if e == nil {
		dbops.UpdateSession(session)
	}
}
