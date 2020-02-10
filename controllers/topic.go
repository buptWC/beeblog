package controllers

import (
	"code.bytedance.com/beeblog/models"
	"code.bytedance.com/beeblog/utils"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = utils.CheckAccountCookie(c.Ctx)
	c.TplName = "topic.html"
	topicList, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error("get all topic failed, err=", err)
	}
	c.Data["Topics"] = topicList
}

func (c *TopicController) Post() {
	if !utils.CheckAccountCookie(c.Ctx) {
		c.Redirect("/", 302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")

	var err error
	tid := c.Input().Get("tid")
	if len(tid) > 0 {
		beego.Info("tid=", tid)
		err = models.ModifyTopic(title, content, tid)
	} else {
		err = models.AddTopic(title, content)
	}
	if err != nil {
		beego.Error("add topic failed, err=", err)
	}
	c.Redirect("/topic", 302)
	return
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}

func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	tid := c.Ctx.Input.Param("0")

	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error("get topic failed, err=", err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
}

func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"
	tid := c.Ctx.Input.Param("0")

	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
}
