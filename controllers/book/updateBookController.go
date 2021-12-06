package book


import (
	"GinWeb/common/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (b *Book) UpdateBook(context *gin.Context) {
	helper.PolarLog.Fatal(context.ClientIP())

	context.JSON(http.StatusOK, gin.H{
		"method": "PUT",
	})
}
