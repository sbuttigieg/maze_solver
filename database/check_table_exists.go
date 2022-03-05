package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func checkIfTableExists(tableName string, db *sql.DB) bool {
	var tableExists bool

	// Query to check if a table exists
	a, err := db.Query(fmt.Sprintf(constants.SQL_CheckTableExists, tableName))
	if err != nil {
		panic(fmt.Sprintf("Failed to check table existence. =>\n%v", err))
	}

	// Scan through the query result for a true or false
	for a.Next() {
		if err := a.Scan(&tableExists); err != nil {
			panic(fmt.Sprintf("Failed to check table existence. =>\n%v", err))
		}
	}
	return tableExists
}
