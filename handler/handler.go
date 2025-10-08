package handler

import (
    "database/sql"

    "github.com/viitorags/encurtadorUrl/config"
)

var (
    logger *config.Logger
    db     *sql.DB
)

func InitializeHandler() {
    logger = config.GetLogger("handler")
    db = config.GetDB()
}
