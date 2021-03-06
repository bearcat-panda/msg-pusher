/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : config.go
#   Created       : 2019/1/7 19:56
#   Last Modified : 2019/1/7 19:56
#   Describe      :
#
# ====================================================*/
package wechat

import (
	"github.com/json-iterator/go"
)

const (
	WeiXinAccessToken = "weixin_access_token"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	APPId     string
	APPSecret string
}
