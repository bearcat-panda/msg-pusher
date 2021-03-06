/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : http.go
#   Created       : 2019/1/8 17:05
#   Last Modified : 2019/1/8 17:05
#   Describe      :
#
# ====================================================*/
package errors

import (
	"bytes"
	"net/http"
	"runtime/debug"

	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ErrHandler 程序奔溃时的返回
func ErrHandler(w http.ResponseWriter, _ *http.Request, err interface{}) {
	handler(w, err)
}

// handler 判断错误时panic还是分装的错误
func handler(w http.ResponseWriter, err interface{}) {
	switch internalErr := err.(type) {
	case *Error:
		logrus.WithFields(logrus.Fields{
			"code":  internalErr.ErrCode,
			"error": internalErr.Error(),
		}).Errorf("Internal error handled")
		toJSON(w, 200, internalErr)
	case error:
		logrus.WithFields(logrus.Fields{
			"error": internalErr,
			"stack": string(debug.Stack()),
		}).Errorf("Internal server error handled")
		toJSON(w, DoErr(internalErr), &Error{
			ErrCode: 1000000,
			Msg:     internalErr.Error(),
			Data:    nil,
		})
	default:
		logrus.WithFields(logrus.Fields{
			"error": err,
			"stack": string(debug.Stack()),
		}).Errorf("Internal server error handled")
		toJSON(w, http.StatusInternalServerError, err)
	}
}

// toJSON 使用json格式返回
func toJSON(w http.ResponseWriter, code int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(buf.Bytes())
}
