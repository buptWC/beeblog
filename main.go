package main

import (
	"code.bytedance.com/beeblog/models"
	_ "code.bytedance.com/beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, false)
	beego.Run()
}
