package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	dns := "root:ganyuefeng1996@tcp(192.168.199.99:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil{
		fmt.Println("main db init failed")
		panic(err)
	}
	router.Static("/static/","./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/",controller.IndexHandle)
	router.GET("/category/",controller.CategoryHandle)
	router.GET("/article/detail/", controller.GetArticledetail)
	router.GET("/article/new/",controller.NewArticle)
	router.POST("/article/submit/", controller.ArticleSubmit)
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.GET("/leave/new/", controller.LeaveNew)
	router.POST("/leave/submit/", controller.LeaveSubmit)
	router.GET("/about/me/", controller.AboutMe)
	router.Run(":8000")
}