package main

import (
	"github.com/gin-gonic/gin"
	"temp/cmd/api"
)

func main() {
	r := gin.Default()
	r.POST("user/login", api.Login)
	r.POST("user/logout", api.Logout)
	r.POST("user/login_status", api.LoginStatus)
	r.POST("user/join_room", api.JoinRoom)
	r.POST("user/step_forward", api.StepForward)
	r.GET("room", api.RoomStatus)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	//var room = svc.NewRoom()
	//var userA = svc.NewUser("a")
	//var userB = svc.NewUser("b")
	//fmt.Println("start")
	//go room.StartGame()
	//go func() {
	//	for {
	//		userA.Run()
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		userB.Run()
	//	}
	//}()
	//fmt.Println(room.GameOver())
}
