package main

import (
	"fmt"

	"gorm.io/driver/mysql" // Change the database driver
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Profile UserProfile // One-to-One relationship
}

type UserProfile struct {
	gorm.Model
	Address string
	UserID  uint `gorm:"uniqueIndex"`
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Referencing User
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// Replace 'user', 'password', 'dbname' with your MySQL credentials and database name

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Auto Migrate
	db.AutoMigrate(&User{}, &UserProfile{})

	// Create User and UserProfile
	user := User{Name: "John"}
	profile := UserProfile{Address: "123 Main St", User: user}

	db.Create(&user)
	db.Create(&profile)

	// Query the user and its profile
	var fetchedUser User
	db.Preload("Profile").First(&fetchedUser, user.ID)

	fmt.Printf("User: %+v\nProfile: %+v\n", fetchedUser, fetchedUser.Profile)
}
