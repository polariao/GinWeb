package main

import (
	"GinWeb/common/helper"
	"GinWeb/models/mysql"
	"GinWeb/route"
	"GinWeb/views"
	"fmt"
	"github.com/gin-gonic/gin"
)



func Init()  {
	//初始化日志
	helper.InitPolarLog("logs")
	//初始化数据库
	mysql.InitConnect()
}

func main() {
	//0、初始化
	Init()
	//1、实例化框架
	r := gin.Default()
	//2、加载渲染模板
	views.InitTemplate(r)
	//3、加载路由
	route.InitRoute(r)
	//4、监听8100端口
	err := r.Run(":8200")
	if nil != err {
		helper.PolarLog.Fatal(fmt.Sprintf("server run fail :%v", err))
	}
}
