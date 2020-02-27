package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "server error"
	c.TplName = "404.html"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "server error"
	c.TplName = "404.html"
}
