package services

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"wx-session/util"
)

type Code2sessionResult struct {
	ErrorCode  int    `json:"errcode"`
	ErrorMsg   string `json:"errmsg,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	ExpiresIn  int    `json:"expires_in,omitempty"`
	Openid     string `json:"openid,omitempty"`
}

type WaterMark struct {
	AppID     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}

type UserInfo struct {
	Openid    string `json:"openid"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	WaterMark `json:"watermark"`
}

var APPID = beego.AppConfig.String("APPID")
var SECRET = beego.AppConfig.String("SECRET")

func Login(code, encrytedData, iv string) (userInfo UserInfo, err error) {
	//get openid and session_key
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + code + "&grant_type=authorization_code"
	r, err := util.HttpGet(url)
	if err != nil {
		//auth error
		fmt.Println(err)
		return
	}
	code2session := Code2sessionResult{}
	fmt.Println(string(r))
	err = json.Unmarshal(r, &code2session)
	if err != nil {
		fmt.Println(err)
		return
	}
	if code2session.ErrorCode > 0 {
		err = fmt.Errorf("%d=>%s", code2session.ErrorCode, code2session.ErrorMsg)
		return
	}
	//code2session success,check signature

	//aes decrypt
	decrypted, err := util.WXBizDataDecrypt(code2session.SessionKey, encrytedData, iv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decrypted))
	//decrypted success,get unionid, weixin unionid is an union id cross multi weixin mp account
	err = json.Unmarshal(decrypted, &userInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	//login success
	//query device by unionid
	return
}
