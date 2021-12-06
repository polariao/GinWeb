package views

import (
	"github.com/gin-gonic/gin"
)

//配置加载模板
func InitTemplate(r *gin.Engine)  {
	//加载静态文件
	r.Static("/public/static","./public/static")
	//加载模板
	r.LoadHTMLGlob("views/**/*")



}
