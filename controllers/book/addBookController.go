package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)
type PostParams struct {
	Name string  `form:"name"`
	Password string `form:"password"`
	File interface{} `form:"file"`
}
func (b *Book) AddBook(context *gin.Context) {
	params := b.PostParams(context)
	fmt.Println(params)
	//文件上传
	file, err := context.FormFile("file")
	if nil != err {
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"文件不存在",
		})
		return
	}
	if  nil != file {
		sprintf :=path.Join("./resources", file.Filename)
		//保存文件
		context.SaveUploadedFile(file,sprintf)
		fmt.Println(sprintf)
	}
	context.JSON(http.StatusOK, gin.H{
		"method": "POST",
		"name":params.Name,
		"password":params.Password,
	})

}

//处理入参
func (b *Book) PostParams(context *gin.Context) PostParams {
	var P PostParams
	err := context.ShouldBind(&P)
	fmt.Println(err)
	return P
}

