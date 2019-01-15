/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : validate.go
#   Created       : 2019/1/15 11:21
#   Last Modified : 2019/1/15 11:21
#   Describe      :
#
# ====================================================*/
package handler

import (
	"strconv"
	"uuabc.com/sendmsg/api/service"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/utils"
)

func checkMobileDetail(mobile, p string) error {
	if !utils.ValidatePhone(mobile) {
		return errors.ErrPhoneNumber
	}
	pg, err := strconv.Atoi(p)
	if err != nil || pg < 1 || pg > 10 {
		return errors.ErrPageInvalidate
	}
	return nil
}

func checkEdit(m service.Meta) error {
	if err := utils.ValidateUUIDV4(m.GetId()); err != nil {
		return err
	}
	return m.ValidateEdit()
}