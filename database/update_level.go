package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func UpdateLevel(id int, value string, db *sql.DB) int {
	var resultingId int

	// Update level
	a, err := db.Query(fmt.Sprintf(constants.SQL_Update_Return_ID,
		Tables["levels"].Name,
		value,
		"id",
		id,
	))
	if err != nil {
		panic(fmt.Sprintf("Failed to update level. =>\n%v", err))
	}

	// Scan through the query result for the id of the updated level
	for a.Next() {
		if err := a.Scan(&resultingId); err != nil {
			panic(fmt.Sprintf("Failed to update level. =>\n%v", err))
		}
	}

	// Return Id of the updated level
	return resultingId
}
