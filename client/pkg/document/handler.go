package document

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func viewDocument(c *gin.Context) {
	documentId := c.Param("documentId")
	zap.L().Debug("will search document in db by this documentId", zap.String("documentId", documentId))
	c.String(http.StatusOK, "view document by documentId"+documentId)
}

func listDocument(c *gin.Context) {

	zap.L().Debug("List all document")
	c.String(http.StatusOK, "List all document")
}
