package v1

import (
	"sync"

	"service/common/components"
	"service/common/model"
)

var userPool = sync.Pool{
	New: func() interface{} {
		return &User{}
	},
}

// NewUser 实例化用户聚合实现
func NewUser() *User {
	return userPool.Get().(*User)
}

// User 用户聚合实现
type User struct{}

func (u *User) RegisterPhone(phone string) (user model.User, err *components.Error) {
	return
}

func (u *User) RegisterEmail(email string) (user model.User, err *components.Error) {
	return
}

func (u *User) Close() error {
	return nil
}
