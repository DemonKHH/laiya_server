package router

import (
	"laiya_server/controllers"
	"laiya_server/middleware"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	r.GET("/api/user/refreshToken", middleware.Authenticate(), controllers.RefreshToken())
	r.GET("/api/getUsers", middleware.Authenticate(), controllers.GetUsers())
	r.GET("/api/getUser", middleware.Authenticate(), controllers.GetUser())
	r.POST("/api/updatePermissions", middleware.Authenticate(), controllers.UpdatePermissions())
	// 需要鉴权的 api
	// auth := r.Group("/auth")
}
