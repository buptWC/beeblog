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
		Title:     categoryName,
		Created:   time.Now(),
		TopicTime: time.Now(),
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

func AddTopic(title, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		beego.Error("insert new topic to orm error, err=", err)
	}
	return err
}

func GetAllTopics(IsReverse bool) (topicList []*Topic, err error) {
	o := orm.NewOrm()
	topicTable := o.QueryTable("topic")

	if IsReverse {
		_, err = topicTable.OrderBy("-created").All(&topicList)
	} else {
		_, err = topicTable.All(&topicList)
	}
	if err != nil {
		beego.Error("get topic from orm error, err=", err)
	}
	return topicList, err
}

func GetTopic(tid string) (topic *Topic, err error) {
	topicId, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("tid is not legal, tid=", tid)
		return topic, err
	}

	o := orm.NewOrm()
	topic = &Topic{
		Id: topicId,
	}

	err = o.Read(topic)
	if err != nil {
		beego.Error("read topic in orm error, err=", err)
		return topic, err
	}

	topic.Views++
	_, err = o.Update(topic)
	if err != nil {
		beego.Error("update topic in orm error, err=", err)
	}
	return topic, err
}

func ModifyTopic(title, content, tid string) error {
	o := orm.NewOrm()
	topic, _ := GetTopic(tid)

	topic.Title = title
	topic.Content = content
	topic.Updated = time.Now()

	_, err := o.Update(topic)
	if err != nil {
		beego.Error("modify topic error, err=", err)
	}
	return err
}

func DeleteTopic(tid string) error {
	topicId, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("tid is not legal, tid=", tid)
		return err
	}

	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	_, err = qs.Filter("id", topicId).Delete()
	if err != nil {
		beego.Error("delete topic error, err=", err)
	}

	return err
}
