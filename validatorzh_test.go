// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validatorzh

import (
	"fmt"
	"strings"
	"testing"
)

type TestData struct {
	username string `validate:"required" label:"用户名"`
	phone    string `validate:"required,mobile" label:"手机号码"`
	idcard   string `validate:"required,idcard" label:"身份证号码"`
	want     string
}

func TestValidate(t *testing.T) {
	var tests = []*TestData{
		{"", "18883339999", "152201197812273005", "用户名必填字段"},
		{"testuser", "1888333999", "152201197812273005", "手机号码格式错误"},
		{"testuser", "18883339999", "15220119781227300", "请输入正确的身份证号码"},
	}

	for key, test := range tests {
		testname := fmt.Sprintf("%d,%s", key, test.want)
		t.Run(testname, func(t *testing.T) {
			err := Validate(&test)
			if strings.EqualFold(test.want, err.Error()) {
				t.Errorf("value is %s,expect is %s", err.Error(), test.want)
			}
		})
	}
}
