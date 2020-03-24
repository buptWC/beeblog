package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddComment(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	comment := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(comment)
	if err != nil {
		beego.Error("insert comment to orm error, err=", err)
		return err
	}
	return nil
}

func GetCommentByTid(tid string) (comments []*Comment, err error) {

	o := orm.NewOrm()
	commentTable := o.QueryTable("comment").Filter("Tid", tid)

	_, err = commentTable.OrderBy("-created").All(&comments)
	if err != nil {
		beego.Error("get comments from orm err, tid=%s, err=%+v", tid, err)
	}
	return comments, err
}

func DeleteCommentById(commentId string) error {
	cid, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		beego.Error("delete comment failed, id not legal, id=", commentId)
		return err
	}

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("id", cid).Delete()
	if err != nil {
		beego.Error("delete topic error, err=", err)
	}
	return err
}
