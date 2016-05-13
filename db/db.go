package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	// postgresql driver
	_ "github.com/lib/pq"
)

var (
	errUnknownDriver = errors.New("Unknown database driver.")
)

// Database struct
type Database struct {
	driver   string
	user     string
	password string
	host     string
	port     string
	database string
}

// NewDatabase return new Database object
func NewDatabase(driver, user, password, host, port, database string) *Database {
	return &Database{driver, user, password, host, port, database}
}

// Connnect and return new connection to database
func (m *Database) Connnect() (*gorm.DB, error) {

	var args string

	// check database driver
	switch m.driver {
	case "mysql":
		args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			m.user, m.password, m.host, m.port, m.database)
	case "postgres":
		args = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			m.user, m.password, m.host, m.port, m.database)
	default:
		return nil, errUnknownDriver
	}

	// open connection
	db, err := gorm.Open(m.driver, args)
	if err != nil {
		return nil, err
	}

	// ping database
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	// logging
	db.LogMode(false)

	return db, nil
}
