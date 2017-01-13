//app.js
var util = require('/utils/util.js')
var qcloud = require('./vendor/qcloud-weapp-client-sdk/index');
var config = require('./config');

App({
  onLaunch: function () {
    qcloud.setLoginUrl(config.service.loginUrl);
    //调用API从本地缓存中获取数据
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)
    //login
    console.log("onLanuch")
    //qcloud.clearSession()
    console.log("start login")
    //qcloud.login()
    console.log("end login")
  },
  globalData:{
    userInfo:null
  }
})