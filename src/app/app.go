package main

import (
	"repository"
	"router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	repository.OpenDB("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	defer repository.CloseDB()
	router.Start()

}
