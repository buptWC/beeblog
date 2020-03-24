package models

import (
	"github.com/astaxie/beego/orm"
	"os"
	"path"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	return err == nil, err
}

func RegisterDB() {
	_, err := PathExists(_DB_NAME)
	if os.IsNotExist(err) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(&Category{}, &Topic{}, &Comment{})
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
