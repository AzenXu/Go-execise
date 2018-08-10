package dbops

import (
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"github.com/gpmgo/gopm/modules/log"
	"database/sql"
)

//  users
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

func QueryUserID(username string) (uid string, e error) {
	stmt, e := db.Prepare(`SELECT id FROM users WHERE login_name = ?`)
	if e != nil {
		log.Error(e.Error())
		return "", e
	}

	e = stmt.QueryRow(username).Scan(&uid); if e != nil {
		log.Error(e.Error())
		return "", e
	}

	return uid, nil
}

func QueryPwd(username string) (pwd string, e error) {
	stmt, e := db.Prepare(`SELECT pwd FROM users WHERE login_name = ?`)
	if e != nil {
		log.Error(e.Error())
		return "", e
	}

	e = stmt.QueryRow(username).Scan(&pwd); if e != nil {
		log.Error(e.Error())
		return "", e
	}

	return pwd, nil
}

//  session
func SelectSession(loginName string) (session *defs.Session, e error) {
	session = &defs.Session{}
	session.UserName = loginName

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

func SelectSessionFromSid(sid string) (session *defs.Session, e error) {
	session = &defs.Session{}
	session.SessionID = sid

	stmt, e := db.Prepare(`SELECT TTL, login_name FROM sessions WHERE session_id = ?`)
	if e != nil {
		log.Error(e.Error())
		return nil, e
	}

	err := stmt.QueryRow(sid).Scan(&session.TTL, &session.UserName)
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

//  videos
func ListAllVideos(username string, from, to int) (videos []*defs.VideoInfo ,e error) {
	stmtOut, e := db.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video_info.create_time DESC`)
	defer stmtOut.Close()

	if e != nil {
		log.Error(e.Error())
		return nil, e
	}

	rows, e := stmtOut.Query(username, from, to); if e != nil {
		log.Error(e.Error())
		return nil, e
	}

	var res []*defs.VideoInfo

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	return res, nil
}