package services

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	// Blank import required to initialize the SQL driver
	_ "github.com/lib/pq"
)

// Postgresql service
type Postgresql struct{}

// Run implements Service interface.
func (Postgresql) Run(parameters map[string]interface{}) (int, error) {
	var (
		dsn string
		err error
		ok  bool
	)

	dsn, ok = parameters["dsn"].(string)
	if !ok || dsn == "" {
		dsn = "postgres://?connect_timeout=5"
	}

	if !strings.Contains(dsn, "connect_timeout") {
		return 1, errors.New(`missing "connect_timeout" parameter in postgresql url`)
	}

	log.Printf(`dsn: "%s"`, dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf(`Error: "%s"`, err)
		return 2, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf(`Error: "%s"`, err)
		return 2, err
	}

	return 0, nil
}

// Name implements Service interface.
func (Postgresql) Name() string {
	return "postgresql"
}

// Parameters implements Service interface.
func (Postgresql) Parameters() []string {
	return []string{
		"dsn",
	}
}
