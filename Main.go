package main

import (
	"Helpin/Controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()
	Router.GET("/Users",Controllers.GetAllPerson)
	Router.GET("/SaveUser",Controllers.SavePerson)
	Router.Run(":8080")
}
