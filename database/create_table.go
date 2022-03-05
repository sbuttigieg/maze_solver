package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func createTable(tableStruct constants.DB_Table, db *sql.DB) {
	// Check if table exists.
	var tableExists bool = checkIfTableExists(tableStruct.Name, db)

	// If table does not exist, create one.
	if tableExists {
		return
	} else {
		_, err := db.Exec(fmt.Sprintf(constants.SQL_TableCreate, tableStruct.Name, tableStruct.SQL_create))
		if err != nil {
			panic(fmt.Sprintf("Failed to create table. =>\n%v", err))
		}
	}

	// Confirm that table was created.
	tableExists = checkIfTableExists(tableStruct.Name, db)
	if !tableExists {
		panic("Table not found after creation.")
	}
}
