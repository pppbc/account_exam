package middleware

import (
	"account_exam/lib/apires"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func ValidToken() gin.HandlerFunc {
	return func(r *gin.Context) {
		cookie, err := r.Cookie("MY_SESSION")

		//验证是否有token
		if err != nil {
			log.Println(err)
			apires.ResFail(r, "no cookie")
			r.Abort()
		} else {
			//TODO 查询redis
			session := sessions.Default(r)
			token := session.Get("MY_SESSION")
			if token == cookie {
				log.Println(cookie)
				r.Next()
			} else {
				log.Println(token)
				log.Printf(cookie)
				apires.ResFail(r, "token wrong")
				r.Abort()
			}
		}
	}
}
