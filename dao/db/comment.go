package db

import (
	"blogger/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InsertComment(comment *model.Comment) (err error)  {
	if comment == nil{
		err = fmt.Errorf("invalid parameter")
		return
	}
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	sqlstr := `insert into comment(content,username,article_id)values(?,?,?)`
	_,err = tx.Exec(sqlstr,comment.Content,comment.Username,comment.ArticleId)
	if err != nil{
		fmt.Println("InserComment failed")
		return
	}
	sqlstr = `  update article set comment_count = comment_count + 1 where id = ?`
	_,err =tx.Exec(sqlstr,comment.ArticleId)
	if err != nil{
		fmt.Println("InserComment update failed")
		return
	}
	err = tx.Commit()
	if err != nil{
		fmt.Println("InserComment tx.Commit failed")
		return
	}
	return
}

func GetCommentList(articleId int64, pageNum int) (commentList []*model.Comment, err error) {
	if pageNum < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d", pageNum)
		return
	}
	sqlstr := `select 
							id, content, username, create_time
						from 
							comment 
						where 
							article_id = ? and 
							status = 1
						order by create_time desc
						`
	//fmt.Println(articleId)
	err = DB.Select(&commentList, sqlstr, articleId)
	if err!= nil{
		fmt.Println("err : ",err)
	}
	fmt.Println("commentList11",len(commentList))
	return
}
