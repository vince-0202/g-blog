package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers(e *gin.Engine) {
	e.GET("/blog", func(c *gin.Context) {
		zap.L().Debug("this is hello func", zap.String("user", "name"), zap.Int("age", 1))
		c.String(http.StatusOK, "hello word")
	})
}
