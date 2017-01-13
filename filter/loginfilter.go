package filter

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
	"wx-session/controllers"
	"wx-session/services"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userinfo").(services.UserInfo)
	fmt.Println(ctx.Request.RequestURI)

	if !ok && !strings.HasSuffix(ctx.Request.RequestURI, "/v1/user/login") {
		ctx.Output.Body([]byte("{\"" + controllers.WX_SESSION_MAGIC_ID + "\":1,\"error\":\"ERR_INVALID_SESSION\"}"))
	}
}
