package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"week_02/app/model"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: user
// @Version: 1.0.0
// @Date: 2020/12/2 20:35
// @Package user

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
}

// UserDao User Dao
type UserDao struct {
}

// GetUser Get User
func (d *UserDao) GetUser(id uint64) (model.User, error) {
	var u model.User
	errScan := db.QueryRow("select accountid,firstName,lastName,"+
		"userno from user where user_no = ? limit 1", id).Scan(&u.AccountID,
		&u.FirstName, &u.LastName, &u.UserNo)
	if errScan != nil {
		return u, errors.Wrap(errScan, "Scan error")
	}

	defer db.Close()
	return u, nil
}
