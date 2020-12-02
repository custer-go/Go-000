package handlers

import (
	"Week02/src/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	uid, status := c.GetQuery("uid")
	if !status {
		c.JSON(http.StatusBadRequest, gin.H{"message": "未传入用户uid"})
		return
	}
	user, err := logic.DefaultUser.FindOne(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "根据uid查找用户信息失败"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": user})
	return
}
