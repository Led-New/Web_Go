package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func BancoDeDados() *sql.DB {
	conexao := "user= postgres dbname = WendelLoja password = 8025w2389 host = localhost sslmode = disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
