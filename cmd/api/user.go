package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temp/cmd/modules"
	"temp/cmd/svc"
)

func Login(c *gin.Context) {
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
			"code":    20001,
			"message": "user already login success",
			"data":    user.(*svc.User).GetName(),
		})
		return
	}
	user := svc.NewUser(u.Name)
	modules.UserPool.Store(u.Name, user)

	c.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "login success",
		"data":    user.GetName(),
	})
	return
}

func Logout(c *gin.Context) {
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
			"message": "logout success",
			"data":    user.(*svc.User).GetName(),
		})
		modules.UserPool.Delete(u.Name)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    30000,
			"message": "logout success",
			"data":    nil,
		})
	}
	return
}

func LoginStatus(c *gin.Context) {
	var u = new(modules.UserLogin)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    30000,
			"message": "user is not login",
			"data":    nil,
		})
		return
	}
	if user, ok := modules.UserPool.Load(u.Name); ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "user is login",
			"data":    user.(*svc.User).GetStatus(),
		})
	}
	return
}

func StepForward(c *gin.Context) {
	var u = new(modules.UserLogin)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    30000,
			"message": "user is not login",
			"data":    nil,
		})
		return
	}
	if user, ok := modules.UserPool.Load(u.Name); ok {
		us := user.(*svc.User)
		if us.GetStatus() == true && us.GetRoom() != nil {
			if us.IsStop() {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    30000,
					"message": "game stop",
					"data":    "ok",
				})
				return
			}
			user.(*svc.User).Run()
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    20000,
				"message": "step success",
				"data":    "ok",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    30000,
				"message": "can not step",
				"data":    "wrong",
			})
			return
		}
	}

}
func JoinRoom(c *gin.Context) {
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
		svc.Join(user.(*svc.User))
		c.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "user is login",
			"data":    "ok",
		})
	}
}
