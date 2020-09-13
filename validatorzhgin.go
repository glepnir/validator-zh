// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validatorzh

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
)

type GinValidatorZh struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *GinValidatorZh) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
	})
}

func (v *GinValidatorZh) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *GinValidatorZh) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		// register mobile
		err := v.validate.RegisterValidation("mobile", mobile)
		if err != nil {
			return err
		}

		// register idcard
		err = v.validate.RegisterValidation("idcard", idcard)
		if err != nil {
			return err
		}

		// register label for better prompt
		v.validate.RegisterTagNameFunc(func(filed reflect.StructField) string {
			name := filed.Tag.Get("label")
			return name
		})

		// i18n
		e := en.New()
		uniTrans := ut.New(e, e, zh.New(), zh_Hant_TW.New())
		translator, _ := uniTrans.GetTranslator("zh")
		zh_translate.RegisterDefaultTranslations(v.validate, translator)

		// 添加手机验证的函数
		v.validate.RegisterTranslation("mobile", translator, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}格式错误", true)
		}, func(ut ut.Translator, ve validator.FieldError) string {
			t, _ := ut.T("mobile", ve.Field(), ve.Field())
			return t
		})

		v.validate.RegisterTranslation("idcard", translator, func(ut ut.Translator) error {
			return ut.Add("idcard", "请输入正确的{0}号码", true)
		}, func(ut ut.Translator, ve validator.FieldError) string {
			t, _ := ut.T("idcard", ve.Field(), ve.Field())
			return t
		})
		var sb strings.Builder

		err = v.validate.Struct(obj)

		if err != nil {
			errs := err.(validator.ValidationErrors)
			for _, err := range errs {
				sb.WriteString(err.Translate(translator))
				sb.WriteString(" ")
			}

			return errors.New(sb.String())
		}
	}

	return nil
}
