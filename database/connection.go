package database

import (
	"github.com/Javohir1649/login/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to databse func
func Connect() {
	conn, err := gorm.Open(mysql.Open("root:Marjonna1649!@/go_login"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = conn

	// shunda databaseda schema ichida go_login fayl ichida users table ochiladi
	//  va ichida id name email boladi
	conn.AutoMigrate(&models.User{})
}
