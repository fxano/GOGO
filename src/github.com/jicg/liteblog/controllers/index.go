package controllers

type IndexController struct {
	BaseController
}
// @router /index [get]
func (this *IndexController) Index() {
	this.TplName="index.html"
}
// @router /about [get]
func (this *IndexController) IndexAbout() {
	this.TplName="about.html"
}
// @router /message [get]
func (this *IndexController) IndexMessage() {
	this.TplName="message.html"
}
//@router /details [get]
func (this *IndexController) IndexDetails(){
	this.TplName="details.html"
}
//@router /comment [get]
func (this *IndexController) IndexComment(){
	this.TplName="comment.html"
}