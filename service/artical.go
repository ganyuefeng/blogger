package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

func GetArticleRecordList(pageNum,pageSize int) (articleRecordList []*model.ArticleRecord,err error) {
	articleInfoList, err := db.GetArticleList(pageNum,pageSize)
	if err!= nil{
		fmt.Println("service GetAllCategoryList failed")
		return nil,err
	}
	if len(articleInfoList)<=0{
		return
	}
	categoryids := GetCategoryIdsbyArticel(articleInfoList)
	categoryList,err := db.GetCategoryList(categoryids)
	if err!= nil{
		fmt.Println("service GetCategoryList failed")
		return nil,err
	}
	for _,articleinfo := range articleInfoList{
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*articleinfo,
		}
		categoryId := articleinfo.Category_id
		for _,category := range categoryList{
			if categoryId == category.CategoryId{
				articleRecord.Category = *category
			}
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}

func GetCategoryIdsbyArticel(articelInfoList []*model.ArticleInfo)(ids []int64)  {
	LABLE:
	for _,articel := range articelInfoList{
		categoryid := articel.Category_id
		for _,id := range ids{
			if id == categoryid{
				continue LABLE
			}
		}
		ids = append(ids,categoryid)
	}
	return
}

func GetArticleRecordListById(categoryId int) (articleRecordList []*model.ArticleRecord,err error) {
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId)
	if err!= nil{
		fmt.Println("service GetAllCategoryList failed")
		return nil,err
	}
	if len(articleInfoList)<=0{
		return
	}
	categoryids := GetCategoryIdsbyArticel(articleInfoList)
	categoryList,err := db.GetCategoryList(categoryids)
	if err!= nil{
		fmt.Println("service GetCategoryList failed")
		return nil,err
	}
	for _,articleinfo := range articleInfoList{
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*articleinfo,
		}
		categoryId := articleinfo.Category_id
		for _,category := range categoryList{
			if categoryId == category.CategoryId{
				articleRecord.Category = *category
			}
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}

func GetArticleDetail(articleid int64)(articleDetail *model.ArticleDetail,err error)  {
	articleDetail,err = db.GetAritcleDetail(articleid)
	if err != nil{
		fmt.Println("GetArticleDetail failed!")
		return
	}
	fmt.Println("category is ",articleDetail.ArticleInfo.Category_id)
	category, err := db.GetCategory(articleDetail.ArticleInfo.Category_id)
	if err != nil{
		fmt.Println("GetArticleDetail GetCategory failed")
		return
	}
	articleDetail.Category = *category
	err = db.UpdateViewCount(articleid)
	if err != nil{
		fmt.Println("UpdateViewCount failed",err)
		return
	}
	return
}

func GetRealtiveArticleList(articleid int64) (realtiveArticleList [] *model.RelativeArticle) {
	realtiveArticleList, err := db.GetRelativeArticle(articleid)
	if err != nil{
		fmt.Println("GetRelativeArticle  failed")
		return
	}
	return
}

func GetPrevAndNextArticleInfo(articleId int64) (prevArticle *model.RelativeArticle,
	nextArticle *model.RelativeArticle, err error) {
	nextArticle, err = db.GetNextArticleInfo(articleId)
	if err != nil{
		fmt.Println("GetNextArticleInfo2  failed")
		return
	}
	if articleId >1 {
		prevArticle, err = db.GetPrevArticleInfo(articleId)
		if err != nil{
			fmt.Println("GetNextArticleInfo2  failed")
			return
		}
	}else{
		prevArticle = &model.RelativeArticle{
			ArticleId: 1,
			Title: "already first",
		}
	}

	return
}

func InsertAricle(article *model.ArticleDetail) (articleId int64,err error) {
	fmt.Println("service article id is ",article.ArticleInfo.Category_id)
	articleId, err = db.InsertAricle(article)
	if err != nil{
		fmt.Println("service InsertAricle failed!")
		return 0, err
	}
	return
}