package homepage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/pkg/environment"
)

func homePage(c *gin.Context) {

	c.HTML(
		http.StatusOK, "homepage.html",
		gin.H{"auth": environment.Env.Author(),
			"info": environment.Env.BlogInfo()},
	)
}
