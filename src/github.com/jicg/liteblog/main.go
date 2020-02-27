package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jicg/liteblog/models"
	_ "github.com/jicg/liteblog/routers"
	"strings"
)

func main() {
	initTemlate()
	beego.Run()

}
func initTemlate(){
	beego.AddFuncMap("equrl",func(x,y string)bool{
		x1:=strings.Trim(x,"/")
		y1:=strings.Trim(y,"/")
		return strings.Compare(x1,y1)==0
	})
}

