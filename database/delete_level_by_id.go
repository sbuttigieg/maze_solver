package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func DeleteLevelById(id int, db *sql.DB) int {
	// Delete level indicated by the id
	result, err := db.Exec(fmt.Sprintf(constants.SQL_DeleteById,
		Tables["levels"].Name,
		"id",
		id,
	))
	if err != nil {
		panic(fmt.Sprintf("Failed to delete level. =>\n%v", err))
	}

	// Return the rows affected
	count, err := result.RowsAffected()
	if err != nil {
		panic(fmt.Sprintf("Failed to delete level. =>\n%v", err))
	}
	return int(count)
}
