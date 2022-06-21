package start

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/pkg/environment"
)

func login(c *gin.Context) {

	if environment.Env.IsFirstDeploy() {
		c.HTML(
			http.StatusOK, "init.html",
			gin.H{"hello": "hello"},
		)
		return
	}

	c.HTML(
		http.StatusOK, "login.html",
		gin.H{"hello": "hello"},
	)
}

func initGblog(c *gin.Context) {

	c.HTML(
		http.StatusOK, "init.html",
		gin.H{"hello": "hello"},
	)
}
