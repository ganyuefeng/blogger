package model

type Category struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	Category_no  int    `db:"category_no"`
}
