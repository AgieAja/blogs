package articlebpc

import (
	articleDac "blogs/dataaccess/articledac"
	helperDac "blogs/dataaccess/helperdac"
	articleModel "blogs/models/articlemodel"
	"log"
)

//GetArticlePublish - get artilce publish
func GetArticlePublish() []articleModel.Articles {
	publish := 3

	myList, err := articleDac.FindByStatus(publish)
	if err != nil {
		log.Println("GetArticlePublish err = " + err.Error())
		return nil
	}

	return myList
}

//GetArticleAll - get all article
func GetArticleAll() []articleModel.VwArticles {
	myList, err := articleDac.FindAll()
	if err != nil {
		log.Println("GetArticleAll err = " + err.Error())
		return nil
	}

	return myList
}

//GetDetailArticle - get detail article
func GetDetailArticle(id int64) articleModel.Articles {
	myData, err := articleDac.FindByArticleID(id)
	if err != nil {
		log.Println("GetDetailArticle err = " + err.Error())
		return myData
	}

	return myData
}

//AddArticle - add new article
func AddArticle(myData articleModel.AddArticle) bool {
	myID, err := helperDac.GetLastID(3)
	if err != nil {
		log.Println("AddArticle err = " + err.Error())
		return false
	}

	errInsert := articleDac.InsertData(myData, myID)
	if errInsert != nil {
		log.Println("AddArticle errInsert = " + errInsert.Error())
		return false
	}

	return true
}

//UpdateArticle - update article
func UpdateArticle(myParam articleModel.AddArticle, myid int64) bool {
	err := articleDac.UpdateData(myParam, myid)
	if err != nil {
		log.Println("UpdateArticle err = " + err.Error())
		return false
	}

	return true
}

//DeleteArticle - delete article
func DeleteArticle(myid int64) bool {
	errDelete := articleDac.DeleteData(myid)
	if errDelete != nil {
		log.Println("DeleteArticle errDelete = " + errDelete.Error())
		return false
	}

	return true
}

//PublishArticle - publish article
func PublishArticle(myid int64) bool {
	err := articleDac.UpdateStatus(myid, 3)
	if err != nil {
		log.Println("PublishArticle err = " + err.Error())
		return false
	}

	return true
}

//UnPublishArticle - unpublish article
func UnPublishArticle(myid int64) bool {
	err := articleDac.UpdateStatus(myid, 2)
	if err != nil {
		log.Println("UnPublishArticle err = " + err.Error())
		return false
	}

	return true
}
