package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/client/pkg/column"
	"github.com/vince-0202/g-blog/client/pkg/document"
	"github.com/vince-0202/g-blog/client/pkg/homepage"
)

type Option func(*gin.Engine)

func (s *Server) InitRoute() {
	s.Include(
		document.Routers,
		homepage.Routers,
		column.Routers,
	)
}
