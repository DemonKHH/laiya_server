package router

import (
	"laiya_server/controllers"
	"laiya_server/middleware"

	"github.com/gin-gonic/gin"
)

func Log(r *gin.Engine) {
	r.POST("/log/upload", middleware.Cors(), controllers.UploadLog())
}
