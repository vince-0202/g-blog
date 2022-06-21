package bootstrap

import (
	"github.com/vince-0202/g-blog/pkg/environment"
)

type GblogServer struct {
	Http HttpServer
	Env  environment.BlogEnv
}
