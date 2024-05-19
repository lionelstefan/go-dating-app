package shared

import (
	"fmt"
	"github.com/lionelstefan/go-dating-app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2/log"
)

func CreateDBConnection() *sqlx.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
		"disable")

	db, err := sqlx.Connect("postgres", conn)

	if err != nil {
		log.Info("Failed connecting to Postgres database")
		log.Error(err.Error())
	} else {
		log.Info("Connected to Postgres database")
	}

	return db
}