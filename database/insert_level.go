package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func InsertNewLevel(db *sql.DB, value string) int {
	var id int

	// Insert new level
	a, err := db.Query(fmt.Sprintf(constants.SQL_Insert_Return_ID,
		Tables["levels"].Name, "level,size_x,size_y,min_path,possible_paths, winning_paths", value))
	if err != nil {
		panic(fmt.Sprintf("Failed to insert new level. =>\n%v", err))
	}

	// Scan through the query result for the id of the inserted level
	for a.Next() {
		if err := a.Scan(&id); err != nil {
			panic(fmt.Sprintf("Failed to insert new level. =>\n%v", err))
		}
	}

	// Return Id of the new level
	return id
}
