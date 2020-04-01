package db

import (
	"database/sql"
	"log"

	"github.com/Dylvn/dragonfire-api/src/config/dotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", dotenv.GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to database.")
}
