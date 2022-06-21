package bootstrap

import (
	"github.com/vince-0202/g-blog/client/pkg/article"
	"github.com/vince-0202/g-blog/client/pkg/column"
	"github.com/vince-0202/g-blog/client/pkg/document"
	"github.com/vince-0202/g-blog/client/pkg/homepage"
)

func (s *Server) InitRoute() {
	s.httpServer.Include(
		document.Routers,
		homepage.Routers,
		column.Routers,
		article.Routers,
	)
}
