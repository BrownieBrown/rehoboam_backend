package main

import (
	"rehoboam/pkg/database"
	"rehoboam/pkg/routes"
	"rehoboam/pkg/utils/env"
)

func main() {
	env.LoadLocalEnvironment()
	database.LoadDatabase()
	routes.SetupRoutes()
}
