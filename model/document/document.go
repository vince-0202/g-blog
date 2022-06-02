package document

import "github.com/vince-0202/g-blog/model/base"

type Document struct {
	base.BaseModel
	Title     string
	Describe  string
	Content   string
	CoverPath string
	ViewNum   int64
}
