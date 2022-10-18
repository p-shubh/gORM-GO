package Connectdatabase

import (
	models "GORM-GO/model"
	"fmt"

	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm@upgrade"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

var DB *gorm.DB

func Connectdatabase() {
	database, err := gorm.Open("sqlite3", "test.db")
	// fmt.Println("databse results", database, err)

	if err != nil {
		panic("Failed to connect to database !")
	}else{
		fmt.Println("db connected succesfully")
	}

	database.AutoMigrate(&models.Books{})

	DB = database
}
