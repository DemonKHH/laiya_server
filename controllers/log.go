package controllers

import (
	response "laiya_server/pkg/common/response"
	"net/http"

	elasticService "laiya_server/service/elastic"

	"github.com/gin-gonic/gin"
)

func UploadLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusOK, response.FailMsg(
				err.Error(),
			))
		}
		c.JSON(http.StatusOK, response.ResponseMsg{
			Code: 0,
			Msg:  "上传成功",
			Data: nil,
		})
		elasticService.Insert(body)
	}
}
