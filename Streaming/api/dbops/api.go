package dbops

import (
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"github.com/gpmgo/gopm/modules/log"
	"database/sql"
)

func Regist(username string, pwd string) (*defs.UserCredential, error) {
	stmt, err := db.Prepare(`INSERT INTO users (login_name, pwd) VALUES (?,?)`)
	defer stmt.Close()

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	_, err = stmt.Exec(username, pwd)
	if err != nil {
		return nil, err
	}

	uc := &defs.UserCredential{Username:username, Pwd:pwd}
	return uc, nil
}

//  session
func SelectSession(loginName string) (session *defs.Session, e error) {
	session = &defs.Session{}
	session.SessionID = loginName

	stmt, e := db.Prepare(`SELECT TTL, session_id FROM sessions WHERE login_name = ?`)
	if e != nil {
		log.Error(e.Error())
		return nil, e
	}

	err := stmt.QueryRow(loginName).Scan(&session.TTL, &session.SessionID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return session, nil
}

func UpdateSession(session defs.Session) (e error) {
	stmt, e := db.Prepare(`UPDATE sessions SET session_id = ?, TTL = ? WHERE login_name = ?`)
	defer stmt.Close()

	if e != nil {
		log.Error(e.Error())
		return e
	}

	_, e = stmt.Exec(session.SessionID, session.TTL, session.UserName)
	if e != nil && e != sql.ErrNoRows {
		log.Error(e.Error())
		return e
	}

	return nil
}

func RegistSession(session defs.Session) (e error) {
	stmt, err := db.Prepare(`INSERT INTO sessions (session_id, login_name, TTL) VALUES (?, ?, ?)`)
	defer stmt.Close()

	if err != nil {
		log.Error(err.Error())
		return err
	}

	_, err = stmt.Exec(session.SessionID, session.UserName, session.TTL)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
