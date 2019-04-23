package articlebpc

import (
	articleDac "blogs/dataaccess/articledac"
	articleModel "blogs/models/articlemodel"
)

//GetArticlePublish - get artilce publish
func GetArticlePublish() []articleModel.Articles {
	publish := 3

	myList, err := articleDac.FindByStatus(publish)
	if err != nil {
		return nil
	}

	return myList
}

//GetArticleAll - get all article
func GetArticleAll() []articleModel.VwArticles {
	myList, err := articleDac.FindAll()
	if err != nil {
		return nil
	}

	return myList
}

//GetDetailArticle - get detail article
func GetDetailArticle(id int64) articleModel.Articles {
	myData, err := articleDac.FindByArticleID(id)
	if err != nil {
		return myData
	}

	return myData
}
