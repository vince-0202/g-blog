package bootstrap

import (
	"github.com/vince-0202/g-blog/admin/pkg/start"
)

func (s *Server) InitRoute() {
	s.httpServer.Include(
		start.Routers,
	)
}
