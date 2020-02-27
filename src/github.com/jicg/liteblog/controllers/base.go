package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}
func (this*BaseController)Prepare(){
	this.Data["Path"]=this.Ctx.Request.RequestURI
}