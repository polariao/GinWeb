package route

import (
	"GinWeb/controllers"
	"GinWeb/controllers/book"
	"GinWeb/controllers/index"
	"fmt"
	"github.com/gin-gonic/gin"
)
//中间件
func autoMiddleWare(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			fmt.Println(123456)
			c.Next()
		}else {
			 c.Next()
		}
	}
}

func InitRoute(r *gin.Engine) {
	//book
	BookGroup := r.Group("book",autoMiddleWare(true))
	{
		BookGroup.GET("/index", (&book.Book{}).GetBook)
		BookGroup.POST("/", (&book.Book{}).AddBook)
		BookGroup.PUT("/", (&book.Book{}).UpdateBook)
		BookGroup.DELETE("/", (&book.Book{}).DeleteBook)
	}


	r.GET("/index", (&index.Index{}).GetIndex)

	r.NoRoute((&controllers.Controller{}).GetIndex)

}
