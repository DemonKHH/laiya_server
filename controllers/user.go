package controllers

import (
	"net/http"

	response "laiya_server/pkg/common/response"
	db "laiya_server/service/mongo"
	serviceUser "laiya_server/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient = db.GetMongoClient()
var userCollection *mongo.Collection = db.OpenCollection(mongoClient, "users")
var validate = validator.New()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// userId := c.Param("userId")
		userId := c.GetString("userId")
		user, err := serviceUser.GetUser(userId)
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.JSON(http.StatusOK, response.FailMsg(
				err.Error(),
			))
		}
		// c.JSON(http.StatusOK, user)
		c.JSON(http.StatusOK, response.ResponseMsg{
			Code: 0,
			Msg:  "获取成功",
			Data: user,
		})
	}
}
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var users []response.LoginResponseWithUser
		users, err = serviceUser.GetUsers()
		if err != nil {
			// c.JSON(http.StatusOK, err.Error())
			c.JSON(http.StatusOK, response.FailMsg(
				err.Error(),
			))
			return
		}
		// c.JSON(http.StatusOK, users)
		c.JSON(http.StatusOK, response.ResponseMsg{
			Code: 0,
			Msg:  "成功获取",
			Data: users,
		})
	}
}

func UpdatePermissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user response.LoginResponseWithUser
		var userInfo struct {
			UserId      string   `json:"userId" binding:"required"`
			Permissions []string `json:"permissions" binding:"required"`
		}
		if err := c.BindJSON(&userInfo); err != nil {
			c.JSON(http.StatusOK, response.FailMsg(
				err.Error(),
			))
			return
		}
		user, err := serviceUser.UpdatePermissions(userInfo.UserId, userInfo.Permissions)
		if err != nil {
			c.JSON(http.StatusOK, response.FailMsg(
				err.Error(),
			))
		}
		c.JSON(http.StatusOK, response.ResponseMsg{
			Code: 0,
			Msg:  "权限更新成功",
			Data: user,
		})
	}
}
