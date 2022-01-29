package v1

import (
	"sync"

	"service/common/define/constant"
	"service/dal/db"
	"service/dal/entity"

	"service/common/components"
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

func (u *User) RegisterPhone(phone string) (user entity.User, err *components.Error) {
	//校验手机号是否已经注册
	info, er := db.UserInfoByPhone(phone)
	if er != nil {
		err = components.NewErrorWithCode(constant.ErrInner)
		return
	}

	if info.Id > 0 {
		err = components.NewErrorWithCode(constant.ErrPhoneExists)
		return
	}

	us := entity.User{
		Phone: phone,
	}

	//添加用户
	er = db.UserAdd(&us)
	if er != nil {
		err = components.NewErrorWithCode(constant.ErrInner)
		return
	}

	//获取刚刚插入的用户信息
	info, er = db.UserInfo(us.Id)
	if er != nil {
		err = components.NewErrorWithCode(constant.ErrInner)
		return
	}

	return info, nil
}

func (u *User) RegisterEmail(email string) (user entity.User, err *components.Error) {
	//校验邮箱是否已经注册
	info, er := db.UserInfoByEmail(email)
	if er != nil {
		err = components.NewErrorWithCode(constant.ErrInner)
		return
	}

	if info.Id > 0 {
		err = components.NewErrorWithCode(constant.ErrEmailExists)
		return
	}

	us := entity.User{
		Email: email,
	}

	//添加用户
	er = db.UserAdd(&us)
	if er != nil {
		err = components.NewErrorWithCode(constant.ErrInner)
		return
	}

	//获取刚刚插入的用户信息
	info, er = db.UserInfo(us.Id)
	if er != nil {
		err = components.NewErrorWithCode(constant.ErrInner)
		return
	}

	return info, nil
}

func (u *User) Close() error {
	return nil
}
