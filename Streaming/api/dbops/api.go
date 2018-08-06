package dbops

import (
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"github.com/gpmgo/gopm/modules/log"
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
