package router

import (
	"laiya_server/controllers"

	"github.com/gin-gonic/gin"
)

func Log(r *gin.Engine) {
	r.POST("/log/upload", controllers.UploadLog())
}
