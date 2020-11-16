package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
	"time"
)

func InsertLeaveMsg(comment, author, email string) (err error) {

	var c model.Leave
	c.Content = comment
	c.Username = author
	c.Email = email
	c.CreateTime = time.Now()
	leavemsgId, err := db.InsertLeaveMsg(&c)
	if err != nil{
		fmt.Println("service InsertLeaveMsg failed",err)
	}
	fmt.Println("InsertLeaveMsg ID is",leavemsgId)
	return
}

func GetLeaveList() (leaveMsgList []*model.Leave, err error) {
	leaveMsgList, err = db.GetLeaveList( 10)
	if err != nil{
		fmt.Println("service GetLeaveList failed",err)
	}
	return
}
