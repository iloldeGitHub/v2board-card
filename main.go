package main

import (
	"github.com/iloldeGitHub/v2board-card/database"
	"github.com/iloldeGitHub/v2board-card/routes"
)

func main() {
	database.InitDB()
	routes.RunRoute()
}
