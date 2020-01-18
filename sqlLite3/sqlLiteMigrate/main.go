package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name string
}
type Profile struct {
	gorm.Model
	User   User `gorm:"foreignkey:UserID"` // use UserID as foreign key
	Name   string
	UserID uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	// db.AutoMigrate(&Product{})
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Profile{})

	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})
	// db.Create(&User{Name: "akif"})

	var profiles []Profile
	// var users []User
	var user User

	db.Where(&User{Name: "ayşe"}).First(&user)
	db.Model(&user).Related(&profiles) //SELECT * from profiles where user_id in ( SELECT id from users) Benzer bir query

	// Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

}
