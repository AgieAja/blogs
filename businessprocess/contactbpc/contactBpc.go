package contactbpc

import (
	contactDac "blogs/dataaccess/contactdac"
	helperDac "blogs/dataaccess/helperdac"

	contactModel "blogs/models/contactmodel"
	"log"
)

//AddMessages - add data messages from guest
func AddMessages(myParam contactModel.AddMessages) bool {
	myID, err := helperDac.GetLastID(1)
	if err != nil {
		log.Println("AddMessages err = " + err.Error())
		return false
	}

	myParam.MessagesID = myID
	errInsert := contactDac.InsertData(myParam)
	if errInsert != nil {
		log.Println("AddMessages errInsert = " + errInsert.Error())
		return false
	}

	return true
}
