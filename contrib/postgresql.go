package contrib

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func CheckPostgresql() error {
	dsn := os.Getenv("DEKITERU_POSTGRESQL_DSN")
	if dsn == "" {
		dsn = "postgres://?connect_timeout=5"
	}
	if !strings.Contains(dsn, "connect_timeout") {
		return errors.New("Missing \"connect_timeout\" parameter in postgresql url.")
	}
	log.Printf("POSTGRESQL_DSN: \"%s\"\n", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Error: \"%s\"\n", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error: \"%s\"\n", err)
		return err
	}
	return nil
}
