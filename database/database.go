package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sbuttigieg/maze_solver/constants"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	// Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s sslmode=%s",
		constants.DB_Host,
		constants.DB_Port,
		constants.DB_User,
		constants.DB_Ssl)

	// Open postgres connection
	DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Printf("Database connection failed. %v\n", err)
		panic(err)
	}

	// Check if database exists.
	dbExists := false
	if checkDatabaseExists(constants.DB_Name, DB) == constants.DB_Name {
		dbExists = true
	}

	// If database exists open it, else create one and open it.
	if !dbExists {
		createDatabase(constants.DB_Name, DB)
	}
	DB = openDatabase(constants.DB_Name)

	// Confirm that the correct database is open.
	currentDB := getCurrentDatabase(DB)
	if currentDB != constants.DB_Name {
		panic("Incorrect database is loaded")
	}

	// Create levels table if it does not exist
	createTable(Tables["levels"], DB)
}
