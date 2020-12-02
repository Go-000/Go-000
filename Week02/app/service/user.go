package service

import (
	"fmt"

	"week_02/app/biz"
	"week_02/app/model"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: user
// @Version: 1.0.0
// @Date: 2020/12/2 20:36
// @Package user

// UserService User Service
type UserService struct {
}

// GetUser Get User
func (u *UserService) GetUser(id uint64) (model.User, error) {
	b := biz.UserBiz{}
	user, err := b.GetUser(id)
	if err != nil {
		return user, fmt.Errorf("service: error occured. \r\n", err)
	}
	fmt.Printf("Found user: %#v\r\n", user)
	return user, err
}
