package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID      uint
	Name    string
	Profile UserProfile
}

type UserProfile struct {
	ID      uint
	Address string
	UserID  uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/coco2?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Auto Migrate
	db.AutoMigrate(&User{}, &UserProfile{})

	// Create User and UserProfile
	user := User{Name: "John", Profile: UserProfile{Address: "123 Main St"}}
	db.Create(&user)

	// Query the user and its profile
	var fetchedUser User
	db.Preload("Profile").First(&fetchedUser, 4)

	// Access the related UserProfile through Association
	var fetchedProfile UserProfile
	db.Model(&fetchedUser).Association("Profile").Find(&fetchedProfile)

	fmt.Printf("User: %+v\nProfile: %+v\n", fetchedUser, fetchedProfile)
}
