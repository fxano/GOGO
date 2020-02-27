package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexComment",
            Router: `/comment`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexDetails",
            Router: `/details`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jicg/liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
