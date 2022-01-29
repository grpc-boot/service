package application

import (
	"service/application/output"
	"service/common/components"
	"service/domain"
)

func getDomain(ver string) (impl interface{}) {
	return domain.V1User()
}

func RegisterPhone(phone string, ver string) (user output.User, err *components.Error) {
	dUser := getDomain(ver).(domain.User)
	defer dUser.Close()

	return output.ConvertUser(dUser.RegisterPhone(phone))
}

func RegisterEmail(email string, ver string) (user output.User, err *components.Error) {
	dUser := getDomain(ver).(domain.User)
	defer dUser.Close()

	return output.ConvertUser(dUser.RegisterEmail(email))
}
