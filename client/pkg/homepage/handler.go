package homepage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", gin.H{"title": "我是测试"})
}
