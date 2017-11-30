package postgres

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	// Blank import required to initialize the SQL driver
	_ "github.com/lib/pq"
)

const (
	defaultDSN = "postgres://?connect_timeout=5"
)

// Run checks PostgreSQL service.
func Run(parameters map[string]interface{}) (int, error) {
	var (
		dsn string
		err error
		ok  bool
	)

	dsn, ok = parameters["dsn"].(string)
	if !ok || dsn == "" {
		dsn = defaultDSN
	}

	if !strings.Contains(dsn, "connect_timeout") {
		return 1, errors.New(`missing "connect_timeout" parameter in postgresql url`)
	}

	log.Printf(`dsn: "%s"\n`, dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf(`Error: "%s"\n`, err)
		return 1, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf(`Error: "%s"\n"`, err)
		return 10, err
	}

	return 0, nil
}
