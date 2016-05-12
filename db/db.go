package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// postgresql driver
	_ "github.com/lib/pq"
)

// Connect to database
func Connect(dbUser, dbPassword, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	options := fmt.Sprintf(`
	user=%s
	password=%s
	host=%s
	port=%s
	dbname=%s
	sslmode=disable`,
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	db, err := gorm.Open("postgres", options)
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.LogMode(false)
	return db, nil
}
