package controllers

import (
	"code.bytedance.com/beeblog/utils"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	if isexit := c.Input().Get("exit") == "true"; isexit {
		c.Ctx.SetCookie("uname", "", -1)
		c.Ctx.SetCookie("pwd", "", -1)
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.Input().Get("username")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"
	// 如果用户名密码通过验证
	if utils.CheckUserAccount(username, pwd) {
		maxAge := 0
		if autoLogin {
			maxAge = (1 << 31) - 1
		}
		c.Ctx.SetCookie("uname", username, maxAge)
		c.Ctx.SetCookie("pwd", pwd, maxAge)
	}

	c.Redirect("/", 302)
	return
}
