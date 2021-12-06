package book

import (
	"GinWeb/common/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
}

type Params struct {
	Name string  `json:"name"`
	Age int `json:"age"`
}
type Response struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Message string `json:"message"`
	Params interface{} `json:"params"`
}

func (b *Book) GetBook(context *gin.Context) {
	helper.PolarLog.Fatal("fatal test 哈哈哈")
	params := b.GetParams(context)
	resp := Response{
		Name: params.Name,
		Age: params.Age,
		Message: "hello word",
	}

	//data := gin.H{
	//	"name": "小王子",
	//	"age": 18,
	//	"message": "hello word",
	//}
	fmt.Println(resp)
	context.JSON(http.StatusOK,resp)
}
//处理入参
func (b *Book) GetParams(context *gin.Context) Params {
	name := context.DefaultQuery("name","不存在")
	age := context.DefaultQuery("age","0")
	ageV1, _ := strconv.Atoi(age)
	return Params{
		Name: name,
		Age: ageV1,
	}

}
