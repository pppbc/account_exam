package lib

import (
	"github.com/gin-gonic/gin"
)

//返回的消息结构
type Info struct {
	Status int         `json:"status"`
	Desc   string      `json:"desc"`
	Data   interface{} `json:"data"`
}

//成功
func ResOk(r *gin.Context, desc string, data interface{}) {
	Res(r, 1, desc, data)
}

//失败
func ResFail(r *gin.Context, desc string) {
	Res(r, 0, desc, nil)
}

//作出响应
func Res(r *gin.Context, status int, desc string, data interface{}) {
	//数据封装
	j := Info{
		Status: status,
		Desc:   desc,
		Data:   data,
	}
	r.JSON(200, j)
}
