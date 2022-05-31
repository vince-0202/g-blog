package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/client/pkg/blog"
	"github.com/vince-0202/g-blog/client/pkg/config"
	"github.com/vince-0202/g-blog/client/pkg/logger"
)

type Option func(*gin.Engine)

type Server struct {
	Server  *gin.Engine
	Options []Option
	Port    int
}

type ClientOptions struct {
	Port int
}

func NewServer(options ClientOptions) (*Server, error) {

	client := &Server{
		Port: options.Port,
	}

	if err := logger.InitLogger(config.DefaultLogConfig()); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return nil, err
	}

	return client, nil
}

func (s *Server) Include(opts ...Option) {
	s.Options = append(s.Options, opts...)
}

func (s *Server) InitSrver() {

	gin.SetMode(gin.ReleaseMode)
	s.Server = gin.New()

	s.Server.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	s.Include(blog.Routers)

	for _, v := range s.Options {
		v(s.Server)
	}
}
