package api

import (
	"week_02/app/model"
	"week_02/app/service"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: user
// @Version: 1.0.0
// @Date: 2020/12/2 20:33
// @Package user

// User user struct
type User struct {
}

// GetUserID Get User ID
func (u *User) GetUserID(id uint64) (model.User, error) {
	s := service.UserService{}
	return s.GetUser(id)
}
