package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func article(c *gin.Context) {

	c.HTML(
		http.StatusOK, "article.html",
		gin.H{"hello": "hello"},
	)
}
