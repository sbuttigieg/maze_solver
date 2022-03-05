package database

import (
	"database/sql"

	"github.com/sbuttigieg/maze_solver/constants"
)

var Tables = map[string]constants.DB_Table{
	"levels": {
		Name:       "levels",
		SQL_create: "id serial primary key, level jsonb, size_x int, size_y int, min_path int, possible_paths int, winning_paths int",
	},
}

// levels table fields
type LevelTableFields_Raw struct {
	Id             int
	Level          []byte
	Size_x         int
	Size_y         int
	Min_path       sql.NullInt16
	Possible_paths sql.NullInt16
	Winning_paths  sql.NullInt16
}

type LevelTableFields struct {
	Id             int
	Level          [][]int
	Size_x         int
	Size_y         int
	Min_path       int
	Possible_paths int
	Winning_paths  int
}
