package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/pkg/environment"
	"github.com/vince-0202/g-blog/pkg/logger"
)

type Server struct {
	Server  *gin.Engine
	Options []Option
	Env     environment.BlogEnv
	Port    int
}

type AdminOptions struct {
	Port int
}

func NewServer(options AdminOptions) (*Server, error) {

	client := &Server{
		Port: options.Port,
		Env:  environment.Env,
	}

	return client, nil
}

func (s *Server) Include(opts ...Option) {
	s.Options = append(s.Options, opts...)
}

func (s *Server) InitSrver() {

	gin.SetMode(gin.ReleaseMode)
	s.Server = gin.New()
	s.Server.LoadHTMLGlob("tem/*")
	s.Server.Static("assets", "./assets")

	s.Server.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	s.initHttpServer()

}

func (s *Server) initHttpServer() {

	s.InitRoute()

	for _, v := range s.Options {
		v(s.Server)
	}

}
