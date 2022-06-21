package start

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/login", login)
	e.GET("/init-gblog", initGblog)
}
