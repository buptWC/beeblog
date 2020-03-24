package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

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
