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
func TestInserComment(t *testing.T) {
	comment := &model.Comment{
		Content: "i love your article!",
		Username: "trmple",
		ArticleId: 3,
	}
	err := InserComment(comment)
	if err != nil{
		fmt.Println("TestInserComment failed")
		return
	}
//}

 */

func TestGetCommentList(t *testing.T) {
	list, err := GetCommentList(3, 10)
	if err != nil{
		fmt.Println("TestGetCommentList failed")
		return
	}
	t.Logf("article1111 : %d\n",len(list))
}