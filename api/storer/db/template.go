/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : template.go
#   Created       : 2019/1/15 15:39
#   Last Modified : 2019/1/15 15:39
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/api/storer"
	"uuabc.com/sendmsg/pkg/pb/tpl"
)

func TemplateInsert(ctx context.Context, templ *tpl.DBTemplate) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO template (id,type,simple_id,content) VALUES (?,?,?,?)")
	if err != nil {
		return tx, err
	}
	_, err = stmt.ExecContext(ctx, templ.Id, templ.Type, templ.SimpleID, templ.Content)
	if err != nil && err.(*mysql.MySQLError).Number == 1062 {
		return tx, ErrUniqueKeyExsits
	}
	return tx, err
}

func TemplateList(ctx context.Context) (res []*tpl.DBTemplate, err error) {
	err = storer.DB.SelectContext(ctx, &res, "SELECT * FROM template")
	return
}