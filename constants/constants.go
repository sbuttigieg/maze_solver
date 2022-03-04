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

type LevelTableFields struct {
	Id             int
	Level          [][]int
	Size_x         int
	Size_y         int
	Min_path       int
	Possible_paths int
	Winning_paths  int
}
