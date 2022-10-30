package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temp/cmd/modules"
	"temp/cmd/svc"
)

func RoomStatus(c *gin.Context) {
	var u = new(modules.UserLogin)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    30000,
			"message": "user msg is invalidate",
			"data":    nil,
		})
		return
	}
	if user, ok := modules.UserPool.Load(u.Name); ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "room is find",
			"data":    user.(*svc.User).GetRoom().GetStatus(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    30000,
		"message": "room is not found",
		"data":    nil,
	})
}
