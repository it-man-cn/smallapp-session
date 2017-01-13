package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"wx-session/services"
)

var WX_SESSION_MAGIC_ID = "F2C224D4-2BCE-4C64-AF9F-A6D872000D1A"

// Operations about Users
type UserController struct {
	beego.Controller
}

type LoginResult struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg,omitempty"`
	Session   `json:"session,omitempty"`
}

type Session struct {
	ID string `json:"id"`
}

// @Title login
// @Description Logs user into the system
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	fmt.Println(u.Ctx.Input.CruSession)
	session := u.Ctx.Input.CruSession
	code := u.Ctx.Input.Header("X-WX-Code")
	encryptedData := u.Ctx.Input.Header("X-WX-Encrypted-Data")
	iv := u.Ctx.Input.Header("X-WX-IV")
	fmt.Println("code:", code)
	fmt.Println("encryptedData:", encryptedData)
	fmt.Println("iv:", iv)
	if len(code) > 0 && len(encryptedData) > 0 && len(iv) > 0 {
		userInfo, err := services.Login(code, encryptedData, iv)
		if err != nil {
			u.Data["json"] = LoginResult{
				ErrorCode: 1,
				ErrorMsg:  "invaliad params",
			}
		} else {
			result := make(map[string]interface{})
			result["session"] = Session{ID: session.SessionID()}
			result[WX_SESSION_MAGIC_ID] = 1
			u.Data["json"] = result
			//save userinfo into session
			session.Set("userinfo", userInfo)
		}
	} else {
		u.Data["json"] = LoginResult{
			ErrorCode: 1,
			ErrorMsg:  "invaliad params",
		}
	}
	u.ServeJSON()
}

// @Title login
// @Description Logs user into the system
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /query [get]
func (u *UserController) Query() {
	session := u.Ctx.Input.CruSession
	if val := session.Get("userinfo"); val != nil {
		u.Data["json"] = val.(services.UserInfo)
	} else {
		u.Data["json"] = LoginResult{
			ErrorCode: 1,
			ErrorMsg:  "need login",
		}
	}
	u.ServeJSON()
}
