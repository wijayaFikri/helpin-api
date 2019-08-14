package Configs

import (
	"Helpin/Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	//open a db connection
	var err error
	Db, err = gorm.Open("mysql", "root:@/HelpinDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")

	}
	//Migrate the schema
	Db.AutoMigrate(&Models.Helper{},
		&Models.User{})
}