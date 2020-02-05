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

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error("add topic failed, err=", err)
	}
	c.Redirect("/topic", 302)
	return
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}
