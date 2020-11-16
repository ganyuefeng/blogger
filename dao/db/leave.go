package db

import (
	"blogger/model"
	"fmt"
)

func InsertLeaveMsg(leavemsg *model.Leave)(leaveId int64,err error)  {
	if len(leavemsg.Content) == 0 {
		fmt.Println("Content is nil")
		return
	}

	sqlstr := "insert into " +
		"leavemsg(content,username,email) value(?,?,?)"

	result,err := DB.Exec(sqlstr,leavemsg.Content,leavemsg.Username,leavemsg.Email)
	if err!= nil{
		fmt.Println("InsertLaveMsg failed")
		return 0,err
	}

	leaveId,err = result.LastInsertId()
	if err!= nil{
		fmt.Println("result get failed")
		return 0,err
	}
	return leaveId,err
}

func GetLeaveList(pageNum int) (leaveMsgList []*model.Leave, err error) {
	if pageNum < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d", pageNum)
		return
	}
	sqlstr := `select 
							id, username, email, content, create_time
						from 
							leavemsg 
						order by create_time desc
						`
	//fmt.Println(articleId)
	err = DB.Select(&leaveMsgList, sqlstr)
	if err!= nil{
		fmt.Println("err : ",err)
	}
	fmt.Println("commentList11",len(leaveMsgList))
	return
}
