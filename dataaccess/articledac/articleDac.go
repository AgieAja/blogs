package articledac

import (
	config "blogs/config"
	articleModel "blogs/models/articlemodel"
	"errors"
)

//FindByStatus - get data from article by status
func FindByStatus(status int) ([]articleModel.Articles, error) {
	db := config.GetMySQLDB()
	if db == nil {
		return nil, errors.New("Failed Connect Database")
	}

	var myList []articleModel.Articles
	err := db.Where("status = ?", status).Find(&myList).Error
	if err != nil {
		return nil, err
	}

	return myList, nil
}

//FindAll - get all data from article
func FindAll() ([]articleModel.VwArticles, error) {
	db := config.GetMySQLDB()
	if db == nil {
		return nil, errors.New("Failed Connect Database")
	}

	var myList []articleModel.VwArticles
	err := db.Find(&myList).Error
	if err != nil {
		return nil, err
	}

	return myList, nil
}

//FindByArticleID - get data from article by article id
func FindByArticleID(id int64) (articleModel.Articles, error) {
	var myData articleModel.Articles
	db := config.GetMySQLDB()
	if db == nil {
		return myData, errors.New("Failed Connect Database")
	}

	err := db.Where("article_id = ?", id).First(&myData).Error
	if err != nil {
		return myData, err
	}

	return myData, nil
}

//InsertData - insert table article
func InsertData(myParam articleModel.AddArticle, myid int64) error {
	db := config.GetMySQLDB()
	if db == nil {
		return errors.New("Failed Connect Database")
	}

	myData := articleModel.Articles{
		Title:            myParam.Title,
		ArticleID:        myid,
		ShortDescription: myParam.ShortDesc,
		Description:      myParam.Desc,
		CreatedBy:        myParam.CreatedBy,
		UpdatedBy:        myParam.UpdatedBy,
		Status:           1,
	}

	err := db.Create(&myData).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateData - update table article
func UpdateData(myParam articleModel.AddArticle, myid int64) error {
	db := config.GetMySQLDB()
	if db == nil {
		return errors.New("Failed Connect Database")
	}

	var myArticle articleModel.Articles
	errUpdate := db.Model(&myArticle).Where("article_id = ?", myid).Updates(articleModel.Articles{
		Title:            myParam.Title,
		ShortDescription: myParam.ShortDesc,
		Description:      myParam.Desc,
		UpdatedBy:        myParam.UpdatedBy,
	}).Error

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

//UpdateStatus - update status table article
func UpdateStatus(myid int64, status int) error {
	db := config.GetMySQLDB()
	if db == nil {
		return errors.New("Failed Connect Database")
	}

	var myArticle articleModel.Articles
	errUpdate := db.Model(&myArticle).Where("article_id = ?", myid).Update(articleModel.Articles{
		Status: status,
	}).Error

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

//DeleteData - delete data table article
func DeleteData(myid int64) error {
	db := config.GetMySQLDB()
	if db == nil {
		return errors.New("Failed Connect Database")
	}

	var myArticle articleModel.Articles
	errDelete := db.Where("article_id = ?", myid).Delete(&myArticle).Error
	if errDelete != nil {
		return errDelete
	}

	return nil
}
