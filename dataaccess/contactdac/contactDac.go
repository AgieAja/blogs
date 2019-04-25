package contactdac

import (
	config "blogs/config"
	contactModel "blogs/models/contactmodel"
	"errors"
)

//InsertData - insert table messages
func InsertData(myParam contactModel.AddMessages) error {
	db := config.GetMySQLDB()
	if db == nil {
		return errors.New("Failed Connect Database")
	}

	myData := contactModel.Messages{
		AddMessages: contactModel.AddMessages{
			MessagesID: myParam.MessagesID,
			FromName:   myParam.FromName,
			FromEmail:  myParam.FromEmail,
			Message:    myParam.Message,
			CreatedBy:  myParam.CreatedBy,
			UpdatedBy:  myParam.UpdatedBy,
		},
	}

	err := db.Create(&myData).Error
	if err != nil {
		return errors.New("err = " + err.Error())
	}

	return nil
}
