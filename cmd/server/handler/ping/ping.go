package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewControllerPing() *Controller {
	return &Controller{}
}

func (c *Controller) HandlerPing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	}
}
