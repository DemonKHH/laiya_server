package main

import (
	ws "laiya_server/lib/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	wsGroup := router.Group("/room")
	{
		wsGroup.GET("/:channel", ws.WebsocketManager.WsClient)
	}
	go ws.WebsocketManager.Start()
	// go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendService()
	router.Run(":3000")
}
