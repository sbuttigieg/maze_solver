package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func getCurrentDatabase(db *sql.DB) string {
	var queryResult string

	// Query to get the current database
	a, err := db.Query(constants.SQL_ShowCurrentDatabase)
	if err != nil {
		panic(fmt.Sprintf("Failed to check name of open database. =>\n%v", err))
	}

	// Scan the query result for the name of the current database
	for a.Next() {
		if err := a.Scan(&queryResult); err != nil {
			panic(fmt.Sprintf("Failed to check name of open database. =>\n%v", err))
		}
	}
	return queryResult
}
