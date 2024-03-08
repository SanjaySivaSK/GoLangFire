package main

import (
	// "fmt"
	// "fmt"
	// "fmt"
	"fmt"
	"log"
	// "slices"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	// gorm.Model
	Id int `gorm:"primaryKey"`

	// Id uint64 `gorm:"primaryKey;autoIncrement:false"`
	Username string
	Email    string
	// TeacherID uint // Foreign key
	// teacher int
}

type teacher struct {
	// gorm.Model
	// Id int `gorm:"primaryKey"`
	name    string
	Student []Student `gorm:"foreignKey:TeacherID"`
}
var student Student
var students []Student

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/GolangTodo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Student{}) // this create table based on the struct name and colunm   based  on field
	// we can able to pass n number struct into the automigrate function   

	// db.AutoMigrate(&teacher{})
     
	// silc:=[]int{2,7,4,6,8,9,5,3}
    // fmt.Println(silc)
	// newSlice:=slices.Delete(silc,1,3)
	// fmt.Println(newSlice)
	

	// Student := Student{"sanjay", "sanjay@282222",}

	// result := db.Create(&Student) 					// pass pointer of data to Create
	// get:=result.RowsAffected      					// we get the row which are affected in the database
	// err=result.Error              					// get error while creating the table and setting 
	// fmt.Println(get,err)

	//db.Select("username").Create(&Student)			// the value will be created only for selected column name  

	// db.Omit("Username").Create(&Student)                // the value will be omitted based of the aregument 

	//Output:

	//  user.Id            // returns inserted data's primary key
	//  result.Error        // returns error
	// result.RowsAffected

	// var user Student
	
	// var student Student
	// records := db.First(&Student)
	// err = records.Error
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// statement := records.Statement
	// if statement != nil  {
	// 	fmt.Println("SQL Statement:", statement)
	// } else {
	// 	fmt.Println("No SQL statement found.")
	// }

	// // Print the retrieved student record
	// fmt.Printf("Retrieved Student: %+v\n", Student)
	

	 db.Find(&students,[]int{1,2})


	// fmt.Println("SQL Statement:", state.SQL.String())
	fmt.Println("Retrieved Student:", students)
}

	
