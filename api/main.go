package main

import (
	"app/models/db"
	"app/router"
)

func main() {
	dbConn := db.ConnectDb()
	defer db.Close()
	router.Router(dbConn)
}
