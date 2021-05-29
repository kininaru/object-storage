package models

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var database *xorm.Engine

func init() {
	db, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		panic(err)
	}
	database, err = xorm.NewEngine("mysql", db)
	if err != nil {
		panic(err)
	}
	err = database.Sync2(new(FileRecord))
	if err != nil {
		panic(err)
	}
	err = database.Sync2(new(User))
	if err != nil {
		panic(err)
	}
}
