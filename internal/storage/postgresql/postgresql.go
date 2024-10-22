package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Connect устанавливает подключение к PostgreSQL с параметрами, переданными из конфигурации
func Connect(user, password, dbname, host string, port int) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Error connecting to PostgreSQL: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("Error pinging PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}
