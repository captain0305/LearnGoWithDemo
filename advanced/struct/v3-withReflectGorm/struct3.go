package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Name  string `gorm:"type:varchar(100);unique_index" validate:"required" format:"lower"`
	Email string `gorm:"type:varchar(100);unique_index" validate:"required,email" format:"lower"`
	Age   int    `gorm:"" validate:"gte=0,lte=130"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Johnny", Email: "johnny@example.com", Age: 20})

	// Validate
	user := &User{Name: "John", Email: "john@email", Age: 200}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}

	// Format
	fmt.Printf("%+v1\n", user)
}
