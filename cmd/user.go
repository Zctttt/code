package cmd

import (
	"time"
)

type IUser interface {
	SetStatus(status bool)
	GetName() string   // 用户名字  前端调用
	GetStep() int      // 获取当前步数
	GetStatus() bool   // 获取当前状态 前端判断是否能摇色子
	SetStep()          // 向下进行一步
	ResetStep(num int) // 触发蛇 梯逻辑
	Run()
}

type User struct {
	status bool
	step   int64
	name   string
	run    chan struct{}
}

func NewUser(name string) *User {
	return &User{
		name:   name,
		step:   0,
		status: false,
		run:    make(chan struct{}),
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
	var t = time.Tick(time.Second)
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
