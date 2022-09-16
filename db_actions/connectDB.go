package db_actions

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:H6Jek0HcJa@tcp(127.0.0.1:3306)/trials_ground?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal("sql.Open error")
	}
	//defer db.Close()

	return db
}
