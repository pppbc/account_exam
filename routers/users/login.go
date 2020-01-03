package users

import (
	"account_exam/lib/apires"
	"account_exam/models"
	"account_exam/proto"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(r *gin.Context) {
	enterpriseId := 12

	//获取用户名密码
	var input proto.LoginUsePasswordInput
	if err := r.Bind(&input); err != nil {
		log.Println(err)
		apires.ResFail(r, "init input failed")
		return
	}

	var user proto.Users
	//通过username获取用户信息
	if err := models.Login.Get(enterpriseId, input.Username, &user); err != nil {
		log.Println(err)
		apires.ResFail(r, "get username failed")
		return
	}
	//验证user是否为空

	//验证密码
	if err := models.Login.ValidPassword(input.Password, user.Password); err != nil {
		apires.ResFail(r, "password wrong")
		return
	}

	//TODO
	//生成token,redis存session,返回cookie

	apires.ResOk(r, "success", user)
}
