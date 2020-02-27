package routers

import (
	"github.com/astaxie/beego"
	"github.com/jicg/liteblog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
    beego.Include(&controllers.IndexController{})
}
