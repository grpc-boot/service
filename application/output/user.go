package output

import (
	"service/infrastructure/utils"
	"service/model"
)

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func ConvertUser(user model.User, err *utils.Error) (u User, er *utils.Error) {
	if err != nil {
		return u, err
	}

	return User{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, err
}
