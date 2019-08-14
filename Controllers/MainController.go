package Controllers

import (
	cf "Helpin/Configs"
	"Helpin/Models"
	"github.com/gin-gonic/gin"
)

var db = cf.Db

func GetAllPerson(context *gin.Context) {
	var users []Models.User
	db.Find(&users)
	context.JSON(200, gin.H{"data": users, "meta": "mblau"})
}

func SavePerson(context *gin.Context) {
	/*	firstname := context.PostForm("FirstName")
		person := Person{FirstName: context.PostForm("FirstName"), LastName: context.PostForm("LastName"), Dob: context.PostForm("dob")}
		db.Save(&person)
		context.JSON(200, gin.H{"data": firstname})*/
	user := Models.User{Firstname: "joko", Lastname: "driyono", Dob: "1996-09-01", Email: "joko_driyono@gov.id", Password: "joyko_choyo", IsSocMed: true, Username: "ah-uh"}
	db.Save(&user)
	context.JSON(200, gin.H{"data": user.Firstname})
}
