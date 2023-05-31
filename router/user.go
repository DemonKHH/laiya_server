package router

import (
	"laiya_server/controllers"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	r.POST("/api/signup", controllers.Signup())
	r.POST("/api/login", controllers.Login())
}
