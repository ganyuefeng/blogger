package model

import "time"

type ArticleInfo struct {
	Id            int64  `db:"id"`
	Category_id   int64  `db:"category_id"`
	Title         string `db:"title"`
	ViewCount    uint32 `db:"view_count"`
	CommentCount uint32 `db:"comment_count"`
	Username      string `db:"username"`
	//Status `db:"status"`
	Summary     string    `db:"summary"`
	CreateTime time.Time `db:"create_time"`
}

type ArticleDetail struct {
	ArticleInfo
	Category
	Content string `db:"content"`
}

type ArticleRecord struct {
	ArticleInfo
	Category
}

type RelativeArticle struct {
	ArticleId int64  `db:"id"`
	Title     string `db:"title"`
}
