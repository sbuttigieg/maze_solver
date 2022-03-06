package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func checkDatabaseExists(dbName string, db *sql.DB) string {
	var databaseExists string

	// Query to check if a database exists
	a, err := db.Query(fmt.Sprintf(constants.SQL_CheckDatabaseExists, dbName))
	if err != nil {
		panic(fmt.Sprintf("Failed to check database existence. =>\n%v", err))
	}

	// Scan through the query result for a true or false
	for a.Next() {
		if err := a.Scan(&databaseExists); err != nil {
			panic(fmt.Sprintf("Failed to check database existence. =>\n%v", err))
		}
	}

	// Returns the database name if exists or an empty string if it doesn't
	return databaseExists
}
