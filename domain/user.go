package domain

import (
	"service/domain/implement/v1"
	"service/infrastructure/abstract"
	"service/infrastructure/utils"
	"service/model"
)

// V1User v1用户聚合实现
func V1User() User {
	return v1.NewUser()
}

// User 用户聚合抽象
type User interface {
	abstract.PoolObject

	RegisterPhone(phone string) (user model.User, err *utils.Error)
	RegisterEmail(email string) (user model.User, err *utils.Error)
}
