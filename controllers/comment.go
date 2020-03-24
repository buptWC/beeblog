package controllers

import (
	"code.bytedance.com/beeblog/models"
	"code.bytedance.com/beeblog/utils"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Get() {

}

func (c *CommentController) Post() {

}

func (c *CommentController) Add() {
	tid := c.Input().Get("tid")
	nickname := c.Input().Get("nickname")
	content := c.Input().Get("content")

	err := models.AddComment(tid, nickname, content)
	if err != nil {
		beego.Error("add comment error, err=", err)
		return
	}

	c.Redirect("/topic/view/"+tid, 302)
	return
}

func (c *CommentController) Delete() {
	if !utils.CheckAccountCookie(c.Ctx) {
		return
	}

	commentId := c.Input().Get("id")
	tid := c.Input().Get("tid")

	err := models.DeleteCommentById(commentId)
	if err != nil {
		beego.Error("delete comment error, commentId=%s, err=%+v", commentId, err)
	}

	c.Redirect("/topic/view/"+tid, 302)
	return
}
