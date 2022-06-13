package environment

import (
	"sync"

	"github.com/vince-0202/g-blog/model/admin"
	"go.uber.org/zap"
)

type BlogConfigEnv struct {
	author         *admin.Author
	blogInfo       *admin.BlogConfig
	InitWatiGroup  *sync.WaitGroup
	InitSuccessTag bool
	FirstDeploy    bool
}

func NewBlogConfigEnv() *BlogConfigEnv {
	b := &BlogConfigEnv{
		InitWatiGroup:  &sync.WaitGroup{},
		InitSuccessTag: false,
		FirstDeploy:    true,
	}

	return b
}

func (bce *BlogConfigEnv) Init() {

	bce.InitWatiGroup.Add(2)

	go bce.initAuthor()
	go bce.initBlogConfig()

	zap.L().Info("waiting for Blog-Config environment init")
	bce.InitWatiGroup.Wait()
	bce.InitSuccessTag = true
	zap.L().Info("Blog-Config environment init success!!!")
}

func (bce *BlogConfigEnv) initAuthor() {
	Env.DB.First(bce.author)
	if bce.author != nil {
		// bce.author = admin.DefaultAuthor()
		zap.L().Info("Ihis G-Blog's author info", zap.String("author name", bce.author.Name))
	}

	bce.InitWatiGroup.Done()

}

func (bce *BlogConfigEnv) initBlogConfig() {
	Env.DB.First(bce.blogInfo)
	if bce.blogInfo != nil {
		// bce.blogInfo = admin.DefaultBlogConfig()
		zap.L().Info("Ihis G-Blog's config info", zap.String("blog name", bce.blogInfo.Name))
	}

	bce.InitWatiGroup.Done()
}
