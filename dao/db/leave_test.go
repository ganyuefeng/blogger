package db

import (
	"blogger/model"
	"fmt"
	"testing"
)

func init()  {
	dns := "root:ganyuefeng1996@tcp(192.168.199.99:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err!= nil{
		panic(err)
	}
	fmt.Println("connect success!")
}

func TestInsertLeaveMsg(t *testing.T) {
	var leavemsg = &model.Leave{
		Username: "hjogdf",
		Email: "ooooo",
		Content: "pppppp",
	}
	leaveId, err := InsertLeaveMsg(leavemsg)
	if err!=nil{
		fmt.Println("TestInsertLeaveMsg failed :",err)
	}
	fmt.Println("leaveId is ",leaveId)
}

func TestGetLeaveList(t *testing.T) {
	_, err := GetLeaveList(5)
	if err!=nil{
		fmt.Println("TestGetLeaveList failed :",err)
	}
}