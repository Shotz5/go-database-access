package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Album struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Artist string
	Price  float32
}

var db *gorm.DB

func main() {
	dsn := "root:password@tcp(localhost:3306)/recordings?charset=utf8mb4"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	db.AutoMigrate(&Album{})

	users := []*Album{
		{
			ID:     1,
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		},
		{
			ID:     2,
			Title:  "Giant Steps",
			Artist: "John Coltrane",
			Price:  63.99,
		},
		{
			ID:     3,
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99,
		},
		{
			ID:     4,
			Title:  "Sarah Vaughan",
			Artist: "Sarah Vaughan",
			Price:  34.98,
		},
	}

	result := db.Create(users)

	if result.Error != nil {
		fmt.Println("Records already created")
	}

	getAllAlbums()
}

func getAllAlbums() {
	var albums []Album
	result := db.Find(&albums)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Albums found: %v\n", albums)
}
