package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func Init(addr string)error  {
	var err error
	DB,err = sqlx.Open("mysql",addr)
	if err != nil{
		fmt.Println("mysql open failed")
		return err
	}
	DB.Ping()
	if err != nil{
		fmt.Println("mysql Ping failed")
		return err
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(50)
	return nil
}