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
	router.GET("/", v1.Index)
	router.GET("/conf", v1.Conf)
	router.GET("/service", v1.Service)
}
