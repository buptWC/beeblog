package routers

import (
	"code.bytedance.com/beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
