package controllers

import (
	"code.bytedance.com/beeblog/models"
	"code.bytedance.com/beeblog/utils"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true // 展示加粗，处于哪个界面
	c.TplName = "home.html"
	c.Data["IsLogin"] = utils.CheckAccountCookie(c.Ctx)

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error("get all topics failed, err=", err)
	}
	c.Data["Topics"] = topics
}
