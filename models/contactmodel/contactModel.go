package contactmodel

import (
	"github.com/jinzhu/gorm"
)

//Messages - struct for table messages
type Messages struct {
	gorm.Model
	AddMessages
}

//AddMessages - struct for add messages
type AddMessages struct {
	MessagesID int64
	FromName   string
	FromEmail  string
	Message    string
	CreatedBy  int64
	UpdatedBy  int64
}
