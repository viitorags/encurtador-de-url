package config

import (
    "database/sql"
    _ "database/sql/driver"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/viitorags/encurtadorUrl/schema"
)

func InitDB() (*sql.DB, error) {
    logger = GetLogger("config")

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        dbUser, dbPass, dbHost, dbPort, dbName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        logger.Error("error initialize database: ", err)
    }

    if err := db.Ping(); err != nil {
        logger.Error("Erro ao conectar no banco: ", err)
        return nil, err
    }

    schema.CreateUrlTable(db)

    return db, nil
}
