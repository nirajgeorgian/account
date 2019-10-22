package db

import (
	log "github.com/Sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/postgres" // import for database
)

// Database :- default global database for CRUD
type Database struct {
	*gorm.DB
}

// New :- database instance for performing actions
func New(config *Config) (*Database, error) {
	db, err := gorm.Open("postgres", config.DatabaseURI)

	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	if err = db.DB().Ping(); err != nil {
		return nil, errors.Wrap(err, "database ping error")
	}
	log.Println("Connected to database")

	return &Database{db}, nil
}
