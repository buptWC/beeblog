package controllers

import "github.com/astaxie/beego"

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
}

func (c *TopicController) Post() {
	c.Ctx.WriteString("post page")
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}
