package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

func GetAllCategoryList()(categoryList []*model.Category,err error)  {
	categoryList, err = db.GetAllCategory()
	if err!= nil{
		fmt.Println("service GetAllCategoryList failed")
		return nil,err
	}
	return
}
