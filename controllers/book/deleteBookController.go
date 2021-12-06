package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (b *Book) DeleteBook(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"method": "DELETE",
	})
}
