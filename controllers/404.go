package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (controller *Controller) GetIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"errno":http.StatusOK,
		"errmsg":"success",
		"data":"本页面被外星人挟持了",
	})
}
