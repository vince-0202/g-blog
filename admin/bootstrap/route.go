package bootstrap

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

func (s *Server) InitRoute() {
	s.Include()
}
