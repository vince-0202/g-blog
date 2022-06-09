package column

import "github.com/vince-0202/g-blog/model/base"

type Column struct {
	base.BaseModel
	Name           string
	CoverPath      string
	Describe       string
	DocumentNum    int
	ShowInHomepage bool
}
