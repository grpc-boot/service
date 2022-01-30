package v1

import (
	"service/application/v1"
	"service/application/v1/output"
	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	phone, _ := ctx.GetQuery("phone")
	email, _ := ctx.GetQuery("email")

	if phone == "" && email == "" {
		json(ctx, components.Code2Response(constant.ErrParam))
		return
	}

	var (
		user output.User
		err  *components.Error
	)

	if phone != "" {
		user, err = v1.RegisterPhone(phone)
	} else if email != "" {
		user, err = v1.RegisterEmail(email)
	}

	if err != nil {
		json(ctx, err.ToResponse())
		return
	}

	json(ctx, model.RespSuccess(user))
}
