package main

import (
	"final_project2/database"
	"final_project2/router"
)

func main() {
	database.MulaiDB()
	r := router.MulaiApp()
	r.Run(":8080")
}
