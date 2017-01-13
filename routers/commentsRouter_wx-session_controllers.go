package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["wx-session/controllers:UserController"] = append(beego.GlobalControllerRouter["wx-session/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["wx-session/controllers:UserController"] = append(beego.GlobalControllerRouter["wx-session/controllers:UserController"],
		beego.ControllerComments{
			"Query",
			`/query`,
			[]string{"get"},
			nil})

}
