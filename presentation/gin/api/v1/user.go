package v1

import (
	"service/application"
	"service/application/output"
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
		user, err = application.RegisterPhone(phone, "v1")
	} else if email != "" {
		user, err = application.RegisterEmail(email, "v1")
	}

	if err != nil {
		json(ctx, err.ToResponse())
		return
	}

	json(ctx, model.RespSuccess(user))
}
