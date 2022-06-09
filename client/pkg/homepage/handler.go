package homepage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vince-0202/g-blog/model/column"
	"github.com/vince-0202/g-blog/pkg/environment"
)

func homePage(c *gin.Context) {

	//list top 3 column by hot

	top3Column := make([]column.Column, 3)
	environment.Env.DB.Where(&column.Column{ShowInHomepage: true}).Find(&top3Column)
	top3Column = append(top3Column, column.Column{Name: "Java", Describe: "This is java column", ShowInHomepage: true})
	top3Column = append(top3Column, column.Column{Name: "Go", Describe: "This is Go column", ShowInHomepage: true})
	top3Column = append(top3Column, column.Column{Name: "Go", CoverPath: "assets/image/upload/Snipaste_2022-05-18_22-08-38.jpg", Describe: "This is Go column", ShowInHomepage: true})

	c.HTML(
		http.StatusOK, "homepage.html",
		gin.H{
			"auth":   environment.Env.Author(),
			"info":   environment.Env.BlogInfo(),
			"column": top3Column,
		},
	)
}
