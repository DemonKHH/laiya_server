package router

import (
	"laiya_server/controllers"
	"laiya_server/middleware"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	r.GET("/user/refreshToken", middleware.Authenticate(), controllers.RefreshToken())
	r.GET("/getUsers", middleware.Authenticate(), controllers.GetUsers())
	r.GET("/getUser", middleware.Authenticate(), controllers.GetUser())
	// 需要鉴权的 api
	// auth := r.Group("/auth")
}
