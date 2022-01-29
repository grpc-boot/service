package db

import (
	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/orm"
)

// 用户db
var userDb = func() orm.Group {
	db, _ := components.GetDb(constant.ContDb)
	return db
}
