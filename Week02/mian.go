package main

import (
	"fmt"

	"week_02/app/api"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: mian.go
// @Version: 1.0.0
// @Date: 2020/12/2 20:32
// @Package Week02

func main() {
	var err error
	var input uint64 = 1
	c := api.User{}
	_, err = c.GetUserID(input)
	fmt.Printf("%+v\r\n", err)
}
