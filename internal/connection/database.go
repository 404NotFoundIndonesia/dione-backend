package connection

import (
	"database/sql"
	"dione-backend/internal/config"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err.Error())
	}

	return db
}
