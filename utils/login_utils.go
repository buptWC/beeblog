package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func CheckUserAccount(username, pwd string) bool {
	return beego.AppConfig.String("uname") == username &&
		beego.AppConfig.String("pwd") == pwd
}

func CheckAccountCookie(ctx *context.Context) bool {
	username := ctx.GetCookie("uname")
	pwd := ctx.GetCookie("pwd")
	return CheckUserAccount(username, pwd)
}
