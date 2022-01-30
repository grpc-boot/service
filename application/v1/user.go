package v1

import (
	"service/application/v1/output"
	"service/common/components"
	"service/domain"
)

func RegisterPhone(phone string) (user output.User, err *components.Error) {
	dUser := domain.V1User().(domain.User)
	defer dUser.Close()

	return output.ConvertUser(dUser.RegisterPhone(phone))
}

func RegisterEmail(email string) (user output.User, err *components.Error) {
	dUser := domain.V1User().(domain.User)
	defer dUser.Close()

	return output.ConvertUser(dUser.RegisterEmail(email))
}
