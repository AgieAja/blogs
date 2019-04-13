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
	CreatedBy        int64
	UpdatedBy        int64
}
