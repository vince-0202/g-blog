package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/client/pkg/blog"
)

type Option func(*gin.Engine)

type Server struct {
	Server   *gin.Engine
	Options  []Option
	Port     int
	LogLevel string
}

type ClientOptions struct {
	Port     int
	LogLevel string
}

func NewServer(options ClientOptions) Server {

	clent := Server{
		Server:   gin.New(),
		Port:     options.Port,
		LogLevel: options.LogLevel,
	}

	return clent
}

func (s *Server) Include(opts ...Option) {
	s.Options = append(s.Options, opts...)
}

func (s *Server) InitSrver() {

	s.Include(blog.Routers)

	for _, v := range s.Options {
		v(s.Server)
	}
}
