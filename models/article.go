package models

import (
	"log"
	"time"
	"web/trials_ground/db_actions"
)

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Tags      []string  `json:"tags"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FgDeleted string    `json:"fg_deleted"`
}

type ArticleList []Article

func GetAll() ArticleList {
	db := db_actions.ConnectDB()
	defer db.Close()
	rows, err := db.Query(db_actions.GetAllArticles)
	if err != nil {
		log.Fatal("sql error GetAll")
	}
	var ls ArticleList
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			log.Fatal("sql error GetAll")
		}
		ls = append(ls, article)
	}
	return ls
}

func GetArticle(id string) Article {
	db := db_actions.ConnectDB()
	defer db.Close()
	var article Article
	err := db.QueryRow(db_actions.GetArticle, id).Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Fatal("sql error GetAll")
	}
	return article
}
