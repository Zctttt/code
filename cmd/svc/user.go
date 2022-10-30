package svc

import (
	"time"
)

type IUser interface {
	GetName() string       // 用户名字  前端调用
	GetStep() int          // 获取当前步数
	GetStatus() bool       // 获取当前状态 前端判断是否能摇色子
	SetStatus(status bool) // 设置状态
	SetStep()              // 向下进行一步
	ResetStep(num int)     // 触发蛇 梯逻辑
	Run()                  // 执行
	GetData() string
	GetRoom() *Room
	IsStop() bool
}

type User struct {
	run    chan struct{}
	status bool
	stop   bool
	step   int64
	room   *Room
	name   string
	data   string
}

func NewUser(name string) *User {
	return &User{
		name:   name,
		step:   0,
		status: false,
		stop:   false,
		run:    make(chan struct{}),
		data:   "",
	}
}

func (u *User) SetStatus(status bool) {
	u.status = status
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetStep() int {
	return int(u.step)
}
func (u *User) GetStatus() bool {
	return u.status
}

func (u *User) SetStep() {
	var t = time.Tick(time.Second * 600)
	var count = 0
	for {
		if count == 5 {
			return
		}
		select {
		case <-u.run:
			u.step += ShackDict()

			return
		case <-t:
			count++
			break
		}
	}

}

func (u *User) ResetStep(num int) {
	u.step = int64(num)
}

func (u *User) Run() {
	u.run <- struct{}{}
}

func (u *User) GetData() string {
	return u.data
}
func (u *User) GetRoom() *Room {
	return u.room
}

func (u *User) IsStop() bool {
	return u.stop
}
