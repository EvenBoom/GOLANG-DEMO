package repository

import (
	"err"
	"struct/entity"
)

//GetPerson is a method to get a person from DB,
//this use prepare that more safer and faster
func GetPerson() *entity.Person {
	sqlDB := ConnDB()
	rows := sqlDB.QueryRow("SELECT * FROM PERSON WHERE `ID` = ?", 1)
	var person *entity.Person
	person = new(entity.Person)
	scanErr := rows.Scan(&person.ID, &person.Job, &person.Name)
	err.PanicError(scanErr)
	return person
}
