package models

import (
	"bytes"
	"io/ioutil"
)

type InstallLogic struct{}

var DefaultInstall = InstallLogic{}

func (InstallLogic) CreateTable() error {
	//objLog := GetLogger(ctx)

	dbFile := "docs/sql/xqj.sql"
	buf, err := ioutil.ReadFile(dbFile)

	if err != nil {
		//objLog.Errorln("create table, read db file error:", err)
		return err
	}

	sqlSlice := bytes.Split(buf, []byte("CREATE TABLE"))
	db.Exec("SET SQL_MODE='ALLOW_INVALID_DATES';")
	for _, oneSql := range sqlSlice {
		strSql := string(bytes.TrimSpace(oneSql))
		if strSql == "" {
			continue
		}

		strSql = "CREATE TABLE " + strSql
		db.Exec(strSql)

	}

	return err
}

// InitTable 初始化数据表
func (InstallLogic) InitTable() error {
	dbFile := "docs/sql/init.sql"
	buf, err := ioutil.ReadFile(dbFile)
	if err != nil {
		//objLog.Errorln("init table, read init file error:", err)
		return err
	}

	sqlSlice := bytes.Split(buf, []byte("INSERT INTO"))
	for _, oneSql := range sqlSlice {
		strSql := string(bytes.TrimSpace(oneSql))
		if strSql == "" {
			continue
		}

		strSql = "INSERT INTO " + strSql
		db.Exec(strSql)
	}
	return err
}
