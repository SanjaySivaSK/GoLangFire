package main

import (
	// "os/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Role struct {
	// Id uint     `gorm:"primarykey"`
	Name string
}
type people struct {
	Id     uint
	Name   string
	Age    int
	RoleId uint `gorm:"foreignKey:RoleID"` // Foreign key
	role   Role
}

func init() {

	// dsn := "root:root@tcp(127.0.0.1:3306)/trial"

	// db, err := gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic("Database Not Connected")

	// }

	// db.AutoMigrate(&Role{}, &people{})
	// adminRole := Role{"Admin"}
	// UserRole := Role{"User"}

	// db.Create(&adminRole)
	// db.Create(&UserRole)

}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/trial"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Database Not Connected")

	}

	db.AutoMigrate(&Role{}, &people{})
	adminRole := Role{"Admin"}
	UserRole := Role{"User"}

	db.Create(&adminRole)
	db.Create(&UserRole)

	// dsn := "root:root@tcp(127.0.0.1:3306)/AppGo"

}
