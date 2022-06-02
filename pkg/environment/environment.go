package environment

import (
	"fmt"

	"github.com/vince-0202/g-blog/model/column"
	"github.com/vince-0202/g-blog/model/document"
	"github.com/vince-0202/g-blog/pkg/config"
	"github.com/vince-0202/g-blog/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Env BlogEnv = BlogEnv{}

type BlogEnv struct {
	Logger *zap.Logger
	Db     *gorm.DB
}

func init() {
	err := Env.initEnv()
	if err != nil {
		fmt.Printf("init g-blog evnironment failed, err:%v\n", err)
	}
}

func (e *BlogEnv) initEnv() error {

	if err := logger.InitLogger(config.DefaultLogConfig()); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return err
	}
	e.Logger = logger.Logger
	zap.L().Info("Initialized g-blog Logger in the BlogEnv!!!")

	e.initDbConnect()
	zap.L().Info("Initialized g-blog DB Conection in the BlogEnv!!!")

	return nil
}

func (e *BlogEnv) initDbConnect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		zap.L().Panic("failed to connect database", zap.Error(err))
	}
	zap.L().Info("Connected to test.db success!!")

	e.Db = db

	err = e.Db.AutoMigrate(&document.Document{}, &column.Column{})
	if err != nil {
		zap.L().Panic("failed to autoMigrate database", zap.Error(err))
	}
}
