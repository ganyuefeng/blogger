package db

import (
	"blogger/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InsertCategory(category *model.Category)  (categoryId int64,err error){
	sqlstr := "insert into category(category_name,category_no) value(?,?)"
	result,err := DB.Exec(sqlstr,category.CategoryName,category.CategoryId)
	if err!= nil{
		fmt.Println("InsertCategory failed")
		return 0,err
	}
	categoryId,err = result.LastInsertId()
	return
}

func GetCategory(id int64) (category *model.Category,err error) {
	category = &model.Category{}
	sqlstr := "select id ,category_name,category_no,category_no from category where id = ?"
	err = DB.Get(category,sqlstr,id)
	if err!= nil{
		fmt.Println("GetCategory failed",err)
		return nil,err
	}
	return
}

func GetCategoryList(categotyIds []int64)(categoryList []*model.Category,err error)  {
	//sqlstr := "select id ,category_name,category_no,category_no from category where id = ?"
	sqlstr,args,err := sqlx.In("select id ,category_name,category_no,category_no from category where id in (?)",categotyIds)
	if err!= nil{
		fmt.Println("GetCategoryList failed",err)
		return nil,err
	}
	err = DB.Select(&categoryList,sqlstr,args...)
	if err!= nil{
		fmt.Println("Select CategoryList failed")
		return nil,err
	}
	return
}

func GetAllCategory() (categoryList []*model.Category,err error) {
//	category = &model.Category{}
	sqlstr := "select id ,category_name,category_no,category_no from category order by category_no asc"
	err = DB.Select(&categoryList,sqlstr)
	if err!= nil{
		fmt.Println("InsertCategory failed")
		return nil,err
	}
	return
}