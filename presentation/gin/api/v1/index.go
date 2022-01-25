package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 20000,
		"msg":  "ok",
		"data": struct{}{},
	})
}
