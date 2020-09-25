// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	validatorzh "github.com/glepnir/validator-zh"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name" validate:"required" lable:"用户名"`
	Phone          string `json:"phone" validate:"required,mobile" label:"联系电话"`
}

func main() {
	e := echo.New()
	e.Validator = &validatorzh.EchoValidatorZh{}
	e.POST("/", func(c echo.Context)error{
		u := new(User)
		if err:=c.Bind(u);err!=nil{
			return err
		}
		return c.String(http.StatusOK, "Ok")
	})
	e.Start(":8080")
}
