package output

import (
	"service/common/components"
	"service/dal/entity"
)

type User struct {
	Id        int64  `json:"id" yaml:"id" xml:"id"`
	Name      string `json:"name" yaml:"name" xml:"name"`
	CreatedAt int64  `json:"created_at" yaml:"created_at" xml:"created_at"`
	UpdatedAt int64  `json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

func ConvertUser(user entity.User, err *components.Error) (u User, er *components.Error) {
	if err != nil {
		return u, err
	}

	return User{
		Id:        user.Id,
		Name:      user.NickName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, err
}
