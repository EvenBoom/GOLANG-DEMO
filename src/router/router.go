package router

import (
	"handler"

	"github.com/gin-gonic/gin"
)

//Start is a method to init router
func Start() {
	router := gin.Default()
	personGroup(router)
	router.Run()
}

//personGroup is a method to get a person message from server
func personGroup(router *gin.Engine) {
	person := router.Group("/person")
	{
		person.GET("/get", handler.GetPerson)
	}

}
