package image

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func loadImage(c *gin.Context) {

	imageId := c.Param("imageId")
	zap.L().Info("load image", zap.String("imageId", imageId))

	path := "image/" + imageId

	file, _ := ioutil.ReadFile(path)
	c.Writer.WriteString(string(file))

}
