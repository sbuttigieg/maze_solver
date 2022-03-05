package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func GetLevelById(id int, db *sql.DB) map[int]LevelTableFields {
	// Map that contains all levels with converted values of Level and Min_path
	level := make(map[int]LevelTableFields)

	// Query to get all levels
	rawLevels, err := db.Query(fmt.Sprintf(constants.SQL_SelectById,
		Tables["levels"].Name,
		"id",
		id,
	))
	if err != nil {
		panic(fmt.Sprintf("Failed to get level by ID. =>\n%v", err))
	}

	// Scan through the query result for values of type LevelTableFields_Raw
	for rawLevels.Next() {
		var queryResultRaw LevelTableFields_Raw
		if err := rawLevels.Scan(
			&queryResultRaw.Id,
			&queryResultRaw.Level,
			&queryResultRaw.Size_x,
			&queryResultRaw.Size_y,
			&queryResultRaw.Min_path,
			&queryResultRaw.Possible_paths,
			&queryResultRaw.Winning_paths,
		); err != nil {
			panic(fmt.Sprintf("Failed to get level by ID. =>\n%v", err))
		}

		// Convert Level into a slice
		var levels [][]int
		err := json.Unmarshal([]byte(string(queryResultRaw.Level[:])), &levels)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal levels. =>\n%v", err))
		}

		// Populate allLevels with the query results (values of Level and Min_path are converted)
		level[queryResultRaw.Id] = LevelTableFields{
			Id:             queryResultRaw.Id,
			Level:          levels,
			Size_x:         queryResultRaw.Size_x,
			Size_y:         queryResultRaw.Size_y,
			Min_path:       int(queryResultRaw.Min_path.Int16),       // converted from int16 to int
			Possible_paths: int(queryResultRaw.Possible_paths.Int16), // converted from int16 to int
			Winning_paths:  int(queryResultRaw.Winning_paths.Int16),  // converted from int16 to int
		}
	}
	return level
}
