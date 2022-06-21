package admin

import "github.com/vince-0202/g-blog/model/base"

type Author struct {
	base.BaseModel
	UserName         string
	Name             string
	ProfilePhotoPath string
	Address          string
	WeChatNum        string
	Email            string
	Motto            string
	Password         string
}

// For Test
func DefaultAuthor() *Author {

	return &Author{
		Name:      "vince",
		Address:   "hangzhou China",
		WeChatNum: "W934778436",
		Email:     "w.vince.0202@gmail.com",
		Motto:     "I can do all things...",
	}
}
