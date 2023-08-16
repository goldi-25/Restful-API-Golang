package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:12345678@tcp(localhost:3306)/mydb?parseTime=true"
var err error

func DatabseMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed:(")
	}
	Database.AutoMigrate(&Employee{})

}
