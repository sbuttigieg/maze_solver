package main

import (
	"fmt"

	"github.com/sbuttigieg/maze_solver/api"
	"github.com/sbuttigieg/maze_solver/constants"
	"github.com/sbuttigieg/maze_solver/database"
)

func main() {
	database.ConnectDB()
	router := api.InitialiseApi()
	router.Run(fmt.Sprintf("%s:%v", constants.Api_Host, constants.Api_Port))
}
