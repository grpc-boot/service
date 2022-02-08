package v1

import (
	"service/application/v1"
	"service/application/v1/output"
	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Register(ctx *routing.Context) error {
	args := ctx.QueryArgs()
	phone := stringArg("phone", args)
	email := stringArg("email", args)

	if phone == "" && email == "" {
		_, err := json(ctx, components.Code2Response(constant.ErrParam))
		return err
	}

	var (
		user output.User
		err  *components.Error
	)

	if len(phone) > 0 {
		user, err = v1.RegisterPhone(phone)
	} else if email != "" {
		user, err = v1.RegisterEmail(email)
	}

	if err != nil {
		_, er := json(ctx, err.ToResponse())
		return er
	}

	_, er := json(ctx, model.RespSuccess(user))
	return er
}
