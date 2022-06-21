package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/bootstrap"
	"github.com/vince-0202/g-blog/pkg/config"
	"github.com/vince-0202/g-blog/pkg/environment"
)

type Server struct {
	httpServer *bootstrap.HttpServer
	Env        environment.BlogEnv
}

type AdminOptions struct {
	Port int
}

func NewServer(options AdminOptions) (*Server, error) {

	httpConfig := config.HttpServerConfig{
		Port: options.Port,
	}

	hs, err := bootstrap.NewHttpServer(httpConfig)
	if err != nil {
		return nil, err
	}

	client := &Server{
		httpServer: hs,
		Env:        environment.Env,
	}

	return client, nil
}

func (s *Server) InitSrver() {

	gin.SetMode(gin.ReleaseMode)
	s.InitRoute()

	s.httpServer.InitHttpServer()

}

func (s *Server) HttpServer() *gin.Engine {
	return s.httpServer.Server
}

func (s *Server) HttpConfig() config.HttpServerConfig {
	return s.httpServer.Config
}
