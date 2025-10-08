package schema

import (
    "database/sql"
    "log"
)

func CreateUrlTable(db *sql.DB) {
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS urls (
        id INT AUTO_INCREMENT PRIMARY KEY,
        original_url TEXT NOT NULL,
        short_url VARCHAR(10) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        access_count INT DEFAULT 0
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `

    _, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatal("Erro ao criar tabela 'urls':", err)
    }
}
