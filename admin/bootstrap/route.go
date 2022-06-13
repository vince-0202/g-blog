package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/admin/pkg/start"
)

type Option func(*gin.Engine)

func (s *Server) InitRoute() {
	s.Include(
		start.Routers,
	)
}
