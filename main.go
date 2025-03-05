package main

import (
	"doc-desc/database"
	"doc-desc/routes"
)

func main() {

	database.InitDatabase()
	r := routes.SetupRouter()

	r.Run(":8080")
}
