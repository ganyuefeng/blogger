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
	fmt.Println("connect success!")
}
/*
func TestGetCategory(t *testing.T)  {
	category,err := GetCategory(1)
	if err!= nil{
		panic(err)
	}
	t.Logf("category :%#v",category)
}

 */

func TestGetListCategory(t *testing.T)  {
	var categoryIds []int64
	categoryIds = append(categoryIds,1,2,3)
	categoryList,err := GetCategoryList(categoryIds)
	if err!= nil{
		panic(err)
	}
	for _,v := range categoryList{
		t.Logf("id:%d category :%#v",v.CategoryId,v)
	}
}
/*
func TestGetALlCategory(t *testing.T)  {
	categoryList,err := GetAllCategory()
	if err!= nil{
		panic(err)
	}
	for _,v := range categoryList{
		t.Logf("id:%d category :%#v",v.CategoryId,v)
	}
}

 */