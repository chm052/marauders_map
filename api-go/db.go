package main

import (
	"fmt"
	_ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

func initDb() *sqlx.DB {
	db_url := ""
	if db_url == "" {
		db_url = "user=marauder password=mArauder dbname=marauder sslmode=disable"
	}
	db, err := sqlx.Open("postgres", db_url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}
