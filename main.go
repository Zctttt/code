package main

import (
	"fmt"
	. "temp/cmd"
)

func main() {
	var room = NewRoom()
	var userA = NewUser("a")
	var userB = NewUser("b")
	_ = room.Join(userA)
	_ = room.Join(userB)
	fmt.Println("start")
	go room.StartGame()
	go func() {
		for {
			userA.Run()
		}
	}()

	go func() {
		for {
			userB.Run()
		}
	}()
	fmt.Println(room.GameOver())
}
