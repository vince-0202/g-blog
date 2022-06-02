package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/client/pkg/blog"
)

type Option func(*gin.Engine)

func (s *Server) InitRoute() {
	s.Include(blog.Routers)
}
