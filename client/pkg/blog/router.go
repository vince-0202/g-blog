package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/blog", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
}
