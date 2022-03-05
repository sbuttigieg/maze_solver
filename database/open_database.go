package database

import (
	"database/sql"
	"fmt"

	"github.com/sbuttigieg/maze_solver/constants"
)

func openDatabase(dbName string) *sql.DB {
	// open the database
	db, err := sql.Open("postgres",
		fmt.Sprintf("dbname=%s sslmode=%s", dbName, constants.DB_Ssl))
	if err != nil {
		panic(fmt.Sprintf("Failed to open database. =>\n%v", err))
	}
	return db
}
