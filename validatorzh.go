// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validatorzh

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(i interface{}) error {
	validate := validator.New()
	// register mobile
	err := validate.RegisterValidation("mobile", mobile)
	if err != nil {
		return err
	}

	// register idcard
	err = validate.RegisterValidation("idcard", idcard)
	if err != nil {
		return err
	}

	// register label for better prompt
	validate.RegisterTagNameFunc(func(filed reflect.StructField) string {
		name := filed.Tag.Get("label")
		return name
	})

	// i18n
	e := en.New()
	uniTrans := ut.New(e, e, zh.New(), zh_Hant_TW.New())
	translator, _ := uniTrans.GetTranslator("zh")
	zh_translate.RegisterDefaultTranslations(validate, translator)

	// 添加手机验证的函数
	validate.RegisterTranslation("mobile", translator, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}格式错误", true)
	}, func(ut ut.Translator, ve validator.FieldError) string {
		t, _ := ut.T("mobile", ve.Field(), ve.Field())
		return t
	})

	validate.RegisterTranslation("idcard", translator, func(ut ut.Translator) error {
		return ut.Add("idcard", "请输入正确的{0}号码", true)
	}, func(ut ut.Translator, ve validator.FieldError) string {
		t, _ := ut.T("idcard", ve.Field(), ve.Field())
		return t
	})
	var sb strings.Builder

	err = validate.Struct(i)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			sb.WriteString(err.Translate(translator))
			sb.WriteString(" ")
		}

		return errors.New(sb.String())
	}
	return nil
}

func mobile(vf validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, vf.Field().String())
	return ok
}

// idcard 验证身份证号码
func idcard(vf validator.FieldLevel) bool {
	id := vf.Field().String()

	var a1Map = map[int]int{
		0:  1,
		1:  0,
		2:  10,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}

	var idStr = strings.ToUpper(string(id))
	var reg, err = regexp.Compile(`^[0-9]{17}[0-9X]$`)
	if err != nil {
		return false
	}
	if !reg.Match([]byte(idStr)) {
		return false
	}
	var sum int
	var signChar = ""
	for index, c := range idStr {
		var i = 18 - index
		if i != 1 {
			if v, err := strconv.Atoi(string(c)); err == nil {
				var weight = int(math.Pow(2, float64(i-1))) % 11
				sum += v * weight
			} else {
				return false
			}
		} else {
			signChar = string(c)
		}
	}
	var a1 = a1Map[sum%11]
	var a1Str = fmt.Sprintf("%d", a1)
	if a1 == 10 {
		a1Str = "X"
	}
	return a1Str == signChar
}
