package admin

import "github.com/vince-0202/g-blog/model/base"

type BlogConfig struct {
	base.BaseModel
	AuthorId  uint
	Name      string
	Greeting  string
	Introduce string
}

// For Test
func DefaultBlogConfig() *BlogConfig {
	return &BlogConfig{
		Name:      "Go-Blog",
		Greeting:  "Welcome to Go Blog",
		Introduce: "Personal Blog build by Golang",
	}
}
