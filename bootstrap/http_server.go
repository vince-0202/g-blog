package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/pkg/config"
	"github.com/vince-0202/g-blog/pkg/logger"
)

type Option func(*gin.Engine)

type HttpServer struct {
	Server  *gin.Engine
	Options []Option
	Config  config.HttpServerConfig
}

func NewHttpServer(options config.HttpServerConfig) (*HttpServer, error) {

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.LoadHTMLGlob("tem/*")
	g.Static("assets", "./assets")

	g.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	httpserver := &HttpServer{
		Config: options,
		Server: g,
	}

	return httpserver, nil
}

func (s *HttpServer) InitHttpServer() {

	for _, v := range s.Options {
		v(s.Server)
	}

}

func (s *HttpServer) Include(opts ...Option) {
	s.Options = append(s.Options, opts...)
}
