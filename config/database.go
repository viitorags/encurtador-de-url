package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/viitorags/encurtadorUrl/schema"
)

func InitDB() (*sql.DB, error) {
	logger = GetLogger("config")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Erro ao inicializar banco de dados: ", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error("Erro ao conectar no banco: ", err)
		return nil, err
	}

	schema.CreateUrlTable(db)

	return db, nil
}
