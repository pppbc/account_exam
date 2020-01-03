package models

import (
	database "account_exam/cmd/db"
	"errors"
	"github.com/jmoiron/sqlx"
)

type LoginController struct {
	db *sqlx.DB
}

var Login = &LoginController{db: database.DB}

func (l LoginController) Get(enterpriseId int, username string, output interface{}) (err error) {
	//noinspection ALL
	query := `
	SELECT * 
	FROM users
	WHERE enterprise_id=$1 AND username=$2
	LIMIT 1
	`
	err = l.db.Get(output, query, enterpriseId, username)
	return
}

func (l LoginController) ValidPassword(userPwd, inputPwd string) (err error) {
	//TODO
	if userPwd != inputPwd {
		err = errors.New("password wrong")
	}
	return
}
