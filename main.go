package main

import (
	"example.com/ece1770/controller/database"
	"example.com/ece1770/router"
)

func main() {
	database.ConnDB()
	r := router.SetupRouter()
	r.Run("0.0.0.0:8080")
}
