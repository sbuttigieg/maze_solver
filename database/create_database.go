package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func createDatabase(dbName string, db *sql.DB) {
	var err error

	// Create database
	_, err = db.Exec(fmt.Sprintf(constants.SQL_CreateDatabase,
		dbName, constants.DB_Enconding, constants.DB_Locale, constants.DB_Locale, constants.DB_Template))
	if err != nil {
		panic(fmt.Sprintf("Failed to create database. =>\n%v", err))
	}
}
