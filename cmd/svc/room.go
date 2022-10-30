package svc

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 多个游戏场景 RoomPool

var EmtyRoom *Room

type Room struct {
	mu        sync.RWMutex
	start     bool
	status    bool
	users     chan *User
	end       chan *User
	userLimit int
	name      string
	Chessboard
}
type Chessboard struct {
	Lattices []*Lattice // 棋盘格子数
}
type Lattice struct {
	index       int
	status      int      // 棋盘状态，0数字，1蛇头，2梯子头，3蛇尾，4梯子尾
	nextLattice *Lattice // 强制跳转到的位置
}

// 新建一个对局
func NewRoom() *Room {
	var room = new(Room)
	room.userLimit = 2
	room.start = false
	room.status = true
	room.end = make(chan *User)
	room.users = make(chan *User, room.userLimit)
	room.name = strconv.Itoa(time.Now().Nanosecond()) // 暂定nanosecond 实际可以用uuid
	room.Lattices = make([]*Lattice, 100)
	for i := 0; i < len(room.Lattices); i++ {
		room.Lattices[i] = &Lattice{
			index:       i + 1,
			status:      0,
			nextLattice: nil,
		}
	}
	room.SetSnake()
	room.SetSnake()
	room.SetSnake()
	room.SetSnake()
	room.SetLadder()
	room.SetLadder()
	room.SetLadder()
	room.SetLadder()
	room.SetSnake()
	room.SetSnake()
	room.SetSnake()
	room.SetSnake()
	room.SetLadder()
	room.SetLadder()
	room.SetLadder()
	room.SetLadder()
	return room
}

// 可能房间人满 默认为2
func Join(user *User) {
	if EmtyRoom == nil {
		EmtyRoom = NewRoom()
	}
	if len(EmtyRoom.users) < EmtyRoom.userLimit {
		user.room = EmtyRoom
		EmtyRoom.users <- user

	}
	if len(EmtyRoom.users) == EmtyRoom.userLimit {
		tmpRoom := EmtyRoom
		EmtyRoom = nil
		go tmpRoom.StartGame()
	}
	return
}

// room在服务其中池化+调度 返回false对局结束，做收尾工作+从池中删除
func (receiver *Room) StartGame() {

	for {
		u := <-receiver.users
		if u.GetStep() != len(receiver.Lattices) && u.GetRoom() != nil {
			u.SetStatus(true)
		}
		var str string
		for {
			if receiver.status == false {
				time.Sleep(time.Second * 600)

			}
			if u.GetStatus() == false {
				receiver.users <- u
				u.data = str
				fmt.Println(str)
				str = ""
				break
			}
			if u.GetStep() == len(receiver.Lattices) {
				u2 := <-receiver.users
				u2.room = nil
				u2.data = "false"
				u2.stop = true
				u.room = nil
				u.data = "win"
				u.stop = true
				receiver.end <- u
				close(receiver.users)
				u.status = false
				return
			}
			u.SetStep()
			u.SetStatus(false)
			str = str + "step " + u.name + strconv.Itoa(u.GetStep())
			if u.GetStep() > len(receiver.Lattices) {
				u.ResetStep(2*len(receiver.Lattices) - u.GetStep())
				str = str + "|back " + u.name + strconv.Itoa(u.GetStep())
			}
			if receiver.Lattices[u.GetStep()-1].status == 1 || receiver.Lattices[u.GetStep()-1].status == 2 {
				u.ResetStep(receiver.Lattices[u.GetStep()-1].nextLattice.index)
				str = str + "|to " + u.name + strconv.Itoa(u.GetStep())
				continue
			}
		}
	}
}

func (receiver *Room) GameOver() *User {
	return <-receiver.end
}

func (receiver *Room) SetSnake() {
	var index = 0
	var head int64 = 0
	var end int64 = 0
	for {
		index++
		if index == 30 {
			return
		}
		head = Rand(2, 99)
		if receiver.Lattices[head].status == 0 {
			break
		}
	}
	for {
		index++
		if index == 30 {
			return
		}
		end = Rand(1, head)
		if receiver.Lattices[end].status == 0 {
			break
		}
	}
	receiver.Lattices[head].status = 1
	receiver.Lattices[head].nextLattice = receiver.Lattices[end]
	receiver.Lattices[end].status = 3

}

func (receiver *Room) SetLadder() {
	var index = 0
	var head int64 = 0
	var end int64 = 0
	for {
		index++
		if index == 30 {
			return
		}
		head = Rand(2, 99)
		if receiver.Lattices[head].status == 0 {
			break
		}
	}
	for {
		index++
		if index == 30 {
			return
		}
		end = Rand(head, 99)
		if receiver.Lattices[end].status == 0 {
			break
		}
	}
	receiver.Lattices[head].status = 2
	receiver.Lattices[head].nextLattice = receiver.Lattices[end]
	receiver.Lattices[end].status = 4

}

func (receiver *Room) GetStatus() bool {
	return receiver.status
}
