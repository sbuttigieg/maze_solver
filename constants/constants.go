package constants

// API Port
const (
	Api_Port = 3000
	Api_Host = "localhost"
)

// Determines the max size in both X and Y directions of the level
const MaxLevelSize int = 100

// List of valid tile types
var LevelObjects = map[string]int{
	"open tile":      0,
	"wall":           1,
	"starting point": 2,
}

// Database constants
const (
	DB_Host      = "localhost"
	DB_Port      = 5432
	DB_User      = "postgres"
	DB_Password  = ""
	DB_Name      = "mazesolver"
	DB_Ssl       = "disable"
	DB_Enconding = "UTF8"
	DB_Locale    = "en_GB"
	DB_Template  = "template0"
)

// Sql queries
var SQL_ListOfDatabases string = "SELECT datname FROM pg_database WHERE datistemplate=false;"
var SQL_CheckDatabaseExists string = "SELECT datname FROM pg_catalog.pg_database WHERE datname = lower('%s')"
var SQL_ShowCurrentDatabase string = "SELECT current_database();"
var SQL_CreateDatabase string = "CREATE DATABASE %s WITH ENCODING '%s'  LC_COLLATE='%s' LC_CTYPE='%s' TEMPLATE=%s;"
var SQL_TableCreate string = "CREATE TABLE %v (%s);"
var SQL_CheckTableExists string = "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema='public' AND table_name='%v');"
var SQL_SelectAll string = "SELECT * FROM %s ORDER BY %v;"
var SQL_SelectById string = "SELECT * FROM %s WHERE %s=%v;"
var SQL_DeleteById string = "DELETE FROM %s WHERE %s=%v;"
var SQL_Insert_Return_ID string = "INSERT INTO %s(%s) values (%s) RETURNING id;"
var SQL_Update_Return_ID string = "UPDATE %s SET %s WHERE %s=%v RETURNING id;"

type DB_Table struct {
	Name       string
	SQL_create string
}
