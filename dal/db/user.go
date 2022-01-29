package db

import (
	"service/dal/entity"

	"github.com/grpc-boot/orm"
)

func UserInfo(id int64) (info entity.User, err error) {
	var (
		where = orm.AndWhere(orm.FieldMap{
			"id": {id},
		})

		user entity.User
	)

	err = userDb().FindOneObj(where, &user, true)

	return user, err
}

func UserInfoByPhone(phone string) (info entity.User, err error) {
	var (
		where = orm.AndWhere(orm.FieldMap{
			"phone": {phone},
		})

		user entity.User
	)

	err = userDb().FindOneObj(where, &user, true)

	return user, err
}

func UserInfoByEmail(email string) (info entity.User, err error) {
	var (
		where = orm.AndWhere(orm.FieldMap{
			"email": {email},
		})

		user entity.User
	)

	err = userDb().FindOneObj(where, &user, true)

	return user, err
}

func UserAdd(user *entity.User) error {
	result, err := userDb().InsertObj(user)
	if err != nil {
		return err
	}

	newId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = newId
	return nil
}
