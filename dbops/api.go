package dbops

import (
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"github.com/gpmgo/gopm/modules/log"
	"database/sql"
	"daker.wang/Azen/Go-execise/Streaming/api/utils"
	"time"
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
	//AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?)
	stmtOut, e := db.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name = ? 
		ORDER BY video_info.create_time DESC`)
	defer stmtOut.Close()

	if e != nil {
		log.Error(e.Error())
		return nil, e
	}

	//rows, e := stmtOut.Query(username, from, to); if e != nil {
	rows, e := stmtOut.Query(username); if e != nil {
		log.Error(e.Error())
		return nil, e
	}

	var res []*defs.VideoInfo

	for rows.Next() {
		var id, name, ctime string
		var aid string
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	log.Warn("%v", res)

	return res, nil
}

func AddNewVideo(aid string, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := db.Prepare(`INSERT INTO video_info 
		(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

//  comments
func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.GenerateUUID()
	if err != nil {
		return err
	}

	stmtIns, err := db.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

//func ListComments(vid string, from, to int) ([]*defs.VideoComment, error) {
//	stmtOut, err := db.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
//		INNER JOIN users ON comments.author_id = users.id
//		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
//		ORDER BY comments.time DESC`)
//
//	var res []*defs.VideoComment
//
//	rows, err := stmtOut.Query(vid, from, to)
//	if err != nil {
//		return res, err
//	}
//
//	for rows.Next() {
//		var id, name, content string
//		if err := rows.Scan(&id, &name, &content); err != nil {
//			return res, err
//		}
//
//		c := &defs.VideoComment{Id: id, VideoId: vid, Author: name, Content: content}
//		res = append(res, c)
//	}
//	defer stmtOut.Close()
//
//	return res, nil
//}