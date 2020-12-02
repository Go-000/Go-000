package biz

import (
	"week_02/app/dao"
	"week_02/app/model"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: user
// @Version: 1.0.0
// @Date: 2020/12/2 20:34
// @Package user

// UserBiz User Biz
type UserBiz struct {
}

// GetUser Get User
func (u *UserBiz) GetUser(id uint64) (model.User, error) {
	d := dao.UserDao{}
	return d.GetUser(id)
}
