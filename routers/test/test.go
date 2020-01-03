package test

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TestLogin(r *gin.Context) {
	session := sessions.Default(r)
	session.Save()
	r.JSON(200, gin.H{"count": 1})
}
