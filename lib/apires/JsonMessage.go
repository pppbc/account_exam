package apires

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回的消息结构
type Info struct {
	Status int         `json:"status"`
	Desc   string      `json:"desc"`
	Data   interface{} `json:"data"`
}

//成功
func ResOk(r *gin.Context, desc string, data interface{}) {
	Res(http.StatusOK, r, 1, desc, data)
}

//接受，但未处理
func ResFail(r *gin.Context, desc string) {
	Res(http.StatusAccepted, r, 0, desc, nil)
}

//作出响应
func Res(code int, r *gin.Context, status int, desc string, data interface{}) {
	//数据封装
	j := Info{
		Status: status,
		Desc:   desc,
		Data:   data,
	}
	//返回
	r.Header("token", "xxxxxxxxxxxxxxxxxxxxxxxxxxx")
	r.JSON(code, j)
}
