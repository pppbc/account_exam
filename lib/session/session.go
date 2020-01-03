package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Creat(session sessions.Session, r *gin.Context, domain int, leftTime time.Duration) {
	session.Set("myss", "123")
	session.Save()
	log.Println("session:", session.Get("myss"))
}
