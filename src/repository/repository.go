package repository

import (
	"database/sql"
	"err"
)

var sqlDB *sql.DB
var dbErr error

//OpenDB is a method to open a DB
func OpenDB(driverName, dataSourceName string) {
	sqlDB, dbErr = sql.Open(driverName, dataSourceName)
	err.PanicError(dbErr)
}

//CloseDB is a method to close a DB
func CloseDB() {
	dbErr = sqlDB.Close()
	err.PanicError(dbErr)
}

//ConnDB is a method to get a DB
func ConnDB() *sql.DB {
	return sqlDB
}
