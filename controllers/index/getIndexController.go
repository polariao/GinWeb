package index

import (
	"GinWeb/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
}

func (b *Index) GetIndex(context *gin.Context) {
	mysql.CreateUser()

	context.JSON(http.StatusOK, gin.H{
		"errno": 0,
		"errmsg": "success",
	})
	//context.HTML(http.StatusOK, "index/index.html", nil)
}
