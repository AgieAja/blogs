package articlemodel

import (
	"github.com/jinzhu/gorm"
)

//Articles - struct for table article
type Articles struct {
	gorm.Model
	ArticleID        int64
	Title            string
	ShortDescription string
	Description      string
	Status           int
	CreatedBy        int64
	UpdatedBy        int64
}

//VwArticles - struct for vw_articles
type VwArticles struct {
	ArticleID        int64
	Title            string
	ShortDescription string
	Description      string
	Status           int
	StatusDesc       string
}

//AddArticle - struct for add new article
type AddArticle struct {
	Title     string
	ShortDesc string
	Desc      string
	CreatedBy int64
	UpdatedBy int64
}
