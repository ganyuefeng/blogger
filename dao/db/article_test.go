package db

import (
	"fmt"
	"testing"
)

func init()  {
	dns := "root:ganyuefeng1996@tcp(192.168.199.99:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err!= nil{
		panic(err)
	}
	fmt.Println("connect123 success!")
}
 /*
func TestInsertAricle(t *testing.T)  {
	article := &model.ArticleDetail{
	}
	article.ArticleInfo.Category_id = 1
	article.ArticleInfo.Comment_count = 0
	article.ArticleInfo.Create_time = time.Now()
	article.ArticleInfo.Title = "test1"
	article.ArticleInfo.Username = "testuser"
	article.ArticleInfo.Summary = "testsummary"
	article.ArticleInfo.View_count = 1
	article.Content = "teststr1111"
	id,err := InsertAricle(article)
	if err != nil{
		fmt.Println("test failed")
	}
	t.Logf("id is %d\n",id)
}





func TestGetArticleList(t *testing.T) {
	articleList,err := GetArticleList(1,15)
	if err != nil{
		return
	}
	//t.Logf("article1111 : %d\n",len(articleList))
}

  */
/*
func TestGetAritcleDetail(t *testing.T) {
	article,err :=GetAritcleDetail(1)
	if err != nil{
		return
	}
	t.Logf("article : %#v\n",article)
}

func TestGetArticleListByCategoryId(t *testing.T)   {
	article,err :=GetArticleListByCategoryId(6,15)
	if err != nil{
		return
	}
	//t.Logf("article6 : %#v\n",article)
	t.Logf("article6 : %#v\n",article)
}



func TestGetRelativeArticle(t *testing.T) {
	article, err := GetRelativeArticle(3)
	if err != nil{
		fmt.Println("TestGetRelativeArticle failed")
		return
	}
	t.Logf("article1111 : %d\n",len(article))
//}

 */
func TestGetNextArticleInfo(t *testing.T) {
	info, err := GetPrevArticleInfo(3)
	if err!= nil{
		fmt.Println("GetPrevArticleInfo failed",err)
		return
	}
	fmt.Println(info)
}