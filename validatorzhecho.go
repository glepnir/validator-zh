// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validatorzh

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
)

type EchoValidatorZh struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *EchoValidatorZh) Validate(i interface{}) error {
	c.lazyInit()

	// register mobile
	err := c.validate.RegisterValidation("mobile", mobile)
	if err != nil {
		return err
	}

	// register idcard
	err = c.validate.RegisterValidation("idcard", idcard)
	if err != nil {
		return err
	}

	// register label for better prompt
	c.validate.RegisterTagNameFunc(func(filed reflect.StructField) string {
		name := filed.Tag.Get("label")
		return name
	})
	// i18n
	e := en.New()
	uniTrans := ut.New(e, e, zh.New(), zh_Hant_TW.New())
	translator, _ := uniTrans.GetTranslator("zh")
	zh_translate.RegisterDefaultTranslations(c.validate, translator)

	// 添加手机验证的函数
	c.validate.RegisterTranslation("mobile", translator, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}格式错误", true)
	}, func(ut ut.Translator, ve validator.FieldError) string {
		t, _ := ut.T("mobile", ve.Field(), ve.Field())
		return t
	})

	c.validate.RegisterTranslation("idcard", translator, func(ut ut.Translator) error {
		return ut.Add("idcard", "请输入正确的{0}号码", true)
	}, func(ut ut.Translator, ve validator.FieldError) string {
		t, _ := ut.T("idcard", ve.Field(), ve.Field())
		return t
	})

	var sb strings.Builder

	err = c.validate.Struct(i)
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

func (c *EchoValidatorZh) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}
