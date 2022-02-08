package core

import (
	v1 "service/presentation/fasthttp/api/v1"

	routing "github.com/qiangxue/fasthttp-routing"
)

// Route 绑定路由
func Route() *routing.Router {
	router := routing.New()

	//绑定路由
	bind(router)

	return router
}

func bind(router *routing.Router) {
	// 加载网关
	router.Use(Gateway)

	router.Get("/", v1.Index)

	//-----------------v1----------------
	v1Group := router.Group("/v1")
	v1Group.Get("/gw", v1.Gw)
	v1Group.Get("/reg", v1.Register)
}
