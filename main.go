package main

import (
	"fmt"

	"github.com/sbuttigieg/maze_solver/api"
	"github.com/sbuttigieg/maze_solver/constants"
)

func main() {
	router := api.InitialiseApi()
	router.Run(fmt.Sprintf("%s:%v", constants.Api_Host, constants.Api_Port))
}
