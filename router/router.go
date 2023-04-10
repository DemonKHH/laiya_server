package router

import (
	"laiya_server/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	log.Printf("[router] load routes")
	Auth(r)
	User(r)
}

func InitRoutes() {
	router := gin.Default()
	router.Use(middleware.Cors())
	LoadRoutes(router)
	router.Run(":8000")
}
