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
