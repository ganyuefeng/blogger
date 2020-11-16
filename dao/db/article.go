package db

import (
	"blogger/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InsertAricle(article *model.ArticleDetail)(articleId int64,err error)  {
	if len(article.Content) == 0 {
		fmt.Println("Content is nil")
		return
	}

	sqlstr := "insert into " +
		"article(content,summary,title,username,category_id,view_count,comment_count) value(?,?,?,?,?,?,?)"

	result,err := DB.Exec(sqlstr,article.Content,article.Summary,article.Title,
		article.Username,article.ArticleInfo.Category_id,article.ViewCount,article.CommentCount)
	if err!= nil{
		fmt.Println("InsertAricle failed")
		return 0,err
	}
	articleId,err = result.LastInsertId()
	if err!= nil{
		fmt.Println("result get failed")
		return 0,err
	}
	return articleId,err
}

func GetArticleList(pageNum,pageSize int) (articlelist [] *model.ArticleInfo,err error) {
	if pageNum<0 || pageSize<0 {
		fmt.Println("pageNum or pageSize is wrong")
		return
	}

	sqlstr := "select id,title,view_count,create_time,comment_count,username,category_id " +
		"from article where status = 1 order by create_time desc limit ?,?"

	err = DB.Select(&articlelist,sqlstr,pageNum,pageSize)
	if err!= nil{
		fmt.Println("GetArticleList get failed",err)
		return nil,err
	}
	return
}

func GetAritcleDetail(articleId int64)(articleDetail *model.ArticleDetail,err error)  {
	if articleId<0 {
		fmt.Println("articleId < 0")
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := "select id,title,view_count,content,create_time,comment_count,username,category_id from article where id = ? and status =1"
	err = DB.Get(articleDetail,sqlstr,articleId)
	fmt.Println("viewcount",articleDetail.ViewCount)
	if err!= nil{
		fmt.Println("GetAritcleDetail get failed",err)
		return nil,err
	}
	return
}

func GetArticleListByCategoryId(categoryId int)(articleList []*model.ArticleInfo,err error) {

	fmt.Println("GetArticleListByCategoryId  :",categoryId)
	sqlstr := "select id,title,view_count,create_time,comment_count,username,category_id " +
		"from article where status = 1 and category_id = ? order by create_time desc "
	err = DB.Select(&articleList,sqlstr,categoryId)
	if err!= nil{
		fmt.Println("GetArticleList get failed",err)
		return nil,err
	}
	return
}

func GetRelativeArticle(articleid int64)  (articlelist []*model.RelativeArticle,err error){
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId,sqlstr,articleid)
	if err!= nil{
		fmt.Println("GetRelativeArticle select category_id failed",err)
		return nil,err
	}

	sqlstr = "select id, title from article where category_id=? and id !=?  limit 10"
	err = DB.Select(&articlelist,sqlstr,categoryId,articleid)
	if err!= nil{
		fmt.Println("GetRelativeArticle select category_id failed",err)
		return nil,err
	}
	return
}

func GetPrevArticleInfo(articleId int64) (articleInfo *model.RelativeArticle,err error) {
	articleInfo = &model.RelativeArticle{
		ArticleId: -1,
	}

	sqlstr := "select id, title from article where id < ? order by id desc limit 1"
	err = DB.Get(articleInfo,sqlstr,articleId)
	if err!= nil{
		fmt.Println("GetPrevArticleInfo failed",err)
		return nil,err
	}

	return
}

func GetNextArticleInfo(articleId int64) (articleInfo *model.RelativeArticle,err error) {
	articleInfo = &model.RelativeArticle{
		ArticleId: -1,
	}

	var maxArticleId int64
	sqlstr := "select max(id) from article"
	err = DB.Get(&maxArticleId,sqlstr)
	if err !=nil{
		fmt.Println("get maxid failed ",err)
	}
	fmt.Println("maxArticleId is ",maxArticleId)

	if articleId==maxArticleId{
		articleInfo = &model.RelativeArticle{
			Title: "already last",
			ArticleId: maxArticleId,
		}
	}else{
		sqlstr := "select id, title from article where id > ? order by id asc limit 1"
		err = DB.Get(articleInfo,sqlstr,articleId)
		if err!= nil{
			fmt.Println("GetNextArticleInfo failed",err)
			return nil,err
		}
	}

	return
}

func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlstr := "select id from article where id=?"
	err = DB.Get(&id, sqlstr, articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}

func UpdateViewCount(articleId int64) (err error) {
	sqlstr := ` update 
						article 
					set 
						view_count = view_count + 1
					where
						id = ?`
	_, err = DB.Exec(sqlstr, articleId)
	if err != nil {
		return
	}
	return
}