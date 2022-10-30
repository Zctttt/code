package modules

import "sync"

var UserPool = sync.Map{}

type UserLogin struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
