package handler

import (
	"repository"

	"github.com/gin-gonic/gin"
)

//GetPerson is a method to get a person message
func GetPerson(c *gin.Context) {
	person := repository.GetPerson()

	c.JSON(200, gin.H{
		"id":   person.ID,
		"name": person.Name,
		"job":  person.Job,
	})

}
