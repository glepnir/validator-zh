// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validatorzh

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

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
