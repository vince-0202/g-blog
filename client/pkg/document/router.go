package document

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/gblog/document/:documentId", viewDocument)
	e.GET("/gblog/document", listDocument)
}
