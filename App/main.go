package main

import (
	

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Role struct {
	// gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"unique"`
	peoples []People
}

type People struct {
	// gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Age      int
	RoleID   uint
	Workouts []Workout
}

type Workout struct {
	ID        uint `gorm:"primaryKey"`
	Workout   string
	Exercises []Exercise
	PeopleID  uint
}

type Exercise struct {
	ID        uint `gorm:"primaryKey"`
	Exercise  string
	WorkoutID uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/coco"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Not Connected")
	}
	db.AutoMigrate(&Role{}, &People{}, &Exercise{}, &Workout{})

	var role []Role
	value := db.Where("name=?", "ADMIN").First(&role)

	if value.RowsAffected == 0 {

		AdminRole := Role{Name: "ADMIN"}
		UserRole := Role{Name: "USER"}
		db.Create(&AdminRole)
		db.Create(&UserRole)

	}
	db.Find(&role)

}
