package cmd

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 多个游戏场景 RoomPool
var RoomPool = sync.Map{}

type Room struct {
	start     bool
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
func (receiver *Room) Join(user *User) error {
	if receiver.start == true {
		return fmt.Errorf("game is start")
	}
	if receiver.userLimit < len(receiver.users) {
		return fmt.Errorf("user number is full")
	}
	receiver.users <- user
	return nil
}

// room在服务其中池化+调度 返回false对局结束，做收尾工作+从池中删除
func (receiver *Room) StartGame() {
	for {
		u := <-receiver.users
		u.SetStatus(true)
		for {
			if u.GetStatus() == false {
				//fmt.Println("ok")
				receiver.users <- u
				break
			}
			if u.GetStep() == len(receiver.Lattices) {
				close(receiver.users)
				receiver.end <- u
				return
			}
			u.SetStep()

			if u.GetStep() > len(receiver.Lattices) {
				u.ResetStep(2*len(receiver.Lattices) - u.GetStep())
			}
			if receiver.Lattices[u.GetStep()-1].status == 1 || receiver.Lattices[u.GetStep()-1].status == 2 {
				u.ResetStep(receiver.Lattices[u.GetStep()-1].nextLattice.index)
				u.SetStatus(false)
				continue
			}
			u.SetStatus(false)
		}
	}
}

func (receiver *Room) GameOver() *User {
	return <-receiver.end
}

func (receiver *Room) SetSnake() {
	var head int64 = 0
	var end int64 = 0
	for {
		head = Rand(2, 99)
		if receiver.Lattices[head].status == 0 {
			break
		}
	}
	for {
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
	var head int64 = 0
	var end int64 = 0
	for {
		head = Rand(2, 99)
		if receiver.Lattices[head].status == 0 {
			break
		}
	}
	for {
		end = Rand(head, 99)
		if receiver.Lattices[end].status == 0 {
			break
		}
	}
	receiver.Lattices[head].status = 2
	receiver.Lattices[head].nextLattice = receiver.Lattices[end]
	receiver.Lattices[end].status = 4

}
