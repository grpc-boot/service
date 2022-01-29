package core

import (
	"service/presentation/gin/api/v1"

	"github.com/gin-gonic/gin"
)

// Route 绑定路由
func Route() *gin.Engine {
	router := gin.Default()

	//绑定路由
	bind(router)

	return router
}

func bind(router *gin.Engine) {
	// 加载网关
	router.Use(WithGateway())

	router.GET("/", v1.Index)

	//-----------------v1----------------
	v1Group := router.Group("v1")
	v1Group.GET("/gw", v1.Gw)
}
