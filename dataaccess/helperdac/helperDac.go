package helperdac

import (
	config "blogs/config"
	helperModel "blogs/models/helpermodel"
	"errors"
)

//GetLastID - get last id from view vw_last_ids
func GetLastID(tableID int64) (int64, error) {
	db := config.GetMySQLDB()
	if db == nil {
		return 0, errors.New("Failed Connect Database")
	}

	var myView helperModel.VwLastIds
	err := db.Where("table_id = ?", tableID).First(&myView).Error
	if err != nil {
		return 0, err
	}

	myID := myView.LastID + 1

	return myID, nil
}
