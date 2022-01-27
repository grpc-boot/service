package application

import (
	"service/application/output"
	"service/common/components"
	"service/domain"
)

type User struct {
	App
}

func (u *User) getDomain(ver string) (impl interface{}) {
	return domain.V1User()
}

func (u *User) RegisterPhone(phone string, ver string) (user output.User, err *components.Error) {
	dUser := u.getDomain(ver).(domain.User)
	defer dUser.Close()

	return output.ConvertUser(dUser.RegisterPhone(phone))
}

func (u *User) RegisterEmail(email string, ver string) (user output.User, err *components.Error) {
	dUser := u.getDomain(ver).(domain.User)
	defer dUser.Close()

	return output.ConvertUser(dUser.RegisterEmail(email))
}
