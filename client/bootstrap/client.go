package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/vince-0202/g-blog/client/pkg/config"
	"github.com/vince-0202/g-blog/client/pkg/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	Server  *gin.Engine
	Options []Option
	Port    int
	Db      *gorm.DB
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

	s.initHttpServer()
	s.initDbConnect()

}

func (s *Server) initHttpServer() {

	s.InitRoute()

	for _, v := range s.Options {
		v(s.Server)
	}

}

func (s *Server) initDbConnect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		zap.L().Panic("failed to connect database")
	}
	zap.L().Info("Connected to test.db success!!")

	s.Db = db
}
