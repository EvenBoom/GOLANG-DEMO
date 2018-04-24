package main

import (
	"repository"
	"router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	repository.OpenDB("mysql", "username:password@tcp(127.0.0.1:3306)/database")
	defer repository.CloseDB()
	router.Start()

}
