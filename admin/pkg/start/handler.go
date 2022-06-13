package start

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {

	c.HTML(
		http.StatusOK, "login.html",
		gin.H{"hello": "hello"},
	)
}
