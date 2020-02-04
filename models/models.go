package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// 分类
type Category struct {
	Id              int64 // 默认是主键
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

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

	orm.RegisterModel(&Category{}, &Topic{})
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(categoryName string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title: categoryName,
	}
	// 先在表中看是否已经有categoryName了
	cateTable := o.QueryTable("category")
	err := cateTable.Filter("title", categoryName).One(cate)
	if err == nil {
		return nil
	}

	_, err = o.Insert(cate)
	if err != nil {
		beego.Error("insert new category to orm error, "+
			"categoryName=", categoryName, "err=", err)
		return err
	}
	return nil
}

func GetAllCategory() (categoryList []*Category, err error) {
	o := orm.NewOrm()

	cateTable := o.QueryTable("category")
	_, err = cateTable.All(&categoryList)
	if err != nil {
		beego.Error("get all category from orm failed, err=", err)
	}
	return categoryList, err
}

func DelCategory(categoryId string) error {
	o := orm.NewOrm()
	categoryId64, err := strconv.ParseInt(categoryId, 10, 64)
	if err != nil {
		beego.Error("categoryId is not legal, categoryId=", categoryId)
		return err
	}
	cate := &Category{
		Id: categoryId64,
	}

	cateTable := o.QueryTable("category")
	err = cateTable.Filter("id", categoryId).One(cate)
	if err != nil {
		beego.Error("categoryId is not exists, categoryId=", categoryId)
		return err
	}

	_, err = cateTable.Filter("id", categoryId).Delete()
	if err != nil {
		beego.Error("delete category by Id failed, categoryId=", categoryId)
		return err
	}
	return nil
}
