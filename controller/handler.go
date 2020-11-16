package controller

import (
	"blogger/model"
	"blogger/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)
func IndexHandle(c *gin.Context)  {
	articleRecordList,err := service.GetArticleRecordList(0,15)
	if err!=nil{
		fmt.Println("controller GetArticleRecordList")
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}

	categoryList,err := service.GetAllCategoryList()
	if err!=nil{
		fmt.Println("controller GetArticleRecordList")
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	c.HTML(http.StatusOK,"views/index.html",gin.H{
		"article_list":articleRecordList,
		"category_list":categoryList,
	})
}

func CategoryHandle(c *gin.Context)  {
	categoryIdStr := c.Query("category_id")
	categoryId,err := strconv.ParseInt(categoryIdStr,10,64)
	if err!=nil{
		fmt.Println("controller GetArticleRecordList")
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	articleRecordList,err := service.GetArticleRecordListById(int(categoryId))
	if err!=nil{
		fmt.Println("controller GetArticleRecordList")
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}

	categoryList, err := service.GetAllCategoryList()
	c.HTML(http.StatusOK,"views/index.html",gin.H{
		"article_list":articleRecordList,
		"category_list":categoryList,
	})
}

func GetArticledetail(c *gin.Context)  {
	articleIdstr := c.Query("article_id")
	articleId,err := strconv.ParseInt(articleIdstr,10,64)
	if err != nil{
		c.HTML(http.StatusInternalServerError,"views/500.html,",nil)
		return
	}
	articleDetail,err := service.GetArticleDetail(articleId)
	if err != nil{
		fmt.Println("handle GetArticleDetail failed")
		c.HTML(http.StatusInternalServerError,"views/500.html,",nil)
		return
	}
	realtiveArticleList := service.GetRealtiveArticleList(articleId)
	if err != nil{
		fmt.Println("handle GetRealtiveArticleList failed")
		c.HTML(http.StatusInternalServerError,"views/500.html,",nil)
		return
	}
	prevArticle, nextArticle, err := service.GetPrevAndNextArticleInfo(articleId)
	if err != nil {
		fmt.Printf("get prev or next article failed, err:%v\n", err)
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil{
		fmt.Println("handle GetAllCategoryList failed")
		c.HTML(http.StatusInternalServerError,"views/500.html,",nil)
		return
	}
	fmt.Println("articleId is ",articleId)
	commentList, err := service.GetCommentList(articleId)
	fmt.Println("commentList is ",len(commentList))
	var m map[string]interface{} = make(map[string]interface{}, 10)
	m["detail"] = articleDetail
	m["relative_article"] = realtiveArticleList
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = allCategoryList
	m["article_id"] = articleId
	m["comment_list"] = commentList
	m["ViewCount"] = articleDetail.ArticleInfo.ViewCount
	fmt.Println("1111",articleDetail.ArticleInfo.ViewCount)
	c.HTML(http.StatusOK, "views/detail.html", m)
}

func NewArticle(c *gin.Context) {
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}
func trimHtml(src string) string {
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}
func ArticleSubmit(c *gin.Context)  {
	content_tmp := c.PostForm("content")
	content := trimHtml(content_tmp)
	fmt.Println("cotent is ",content)
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")
	title := c.PostForm("title")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.ArticleInfo.Category_id = categoryId
	articleDetail.Username = author
	articleDetail.Title = title
	fmt.Printf("handle title = %s,Category_id = %d\n",articleDetail.Title,articleDetail.Category_id)
	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])
	articleId, err := service.InsertAricle(articleDetail)
	if err!= nil{
		fmt.Println("service insert article failed")
	}
	fmt.Printf("insert article succ, id:%d, err:%v\n", articleId, err)
	c.Redirect(http.StatusMovedPermanently, "/")
	return
}

func CommentSubmit(c *gin.Context)  {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	articleIdstr := c.PostForm("article_id")
	fmt.Println(articleIdstr)
	articleId,err := strconv.ParseInt(articleIdstr,10,64)
	if err !=nil{
		fmt.Println("CommentSubmit parseint failed")
		return
	}
	err = service.InsertComment(comment,author,email,articleId)
	if err != nil {
		fmt.Printf("insert comment failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d", articleId)
	c.Redirect(http.StatusMovedPermanently, url)
}

func LeaveNew(c *gin.Context) {
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/gbook.html", leaveList)
}


func LeaveSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	fmt.Println(comment,author,email)
	err := service.InsertLeaveMsg(author, email, comment)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/leave/new/")
	c.Redirect(http.StatusMovedPermanently, url)
}

func AboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "views/about.html", gin.H{
		"title": "Posts",
	})
}