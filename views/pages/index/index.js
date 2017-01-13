//index.js
//获取应用实例
var util = require('../../utils/util.js')
var qcloud = require('../../vendor/qcloud-weapp-client-sdk/index');
var app = getApp()
Page({
  data: {
    motto: 'Hello World',
    userInfo: {}
  },
  //事件处理函数
  bindViewTap: function() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  onLoad: function () {
    console.log('onLoad')
    var that = this
    qcloud.request({
      login:true,
      url:"https://cb.tunnel.litelink.me/v1/user/query",
      success:function(res){
        console.log(res.data)
        that.setData({
           userInfo:res.data
        })
      }
    })
  }
})
