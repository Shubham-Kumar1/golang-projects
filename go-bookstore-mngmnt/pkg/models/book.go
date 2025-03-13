package models

import (
	"github.com/Shubham-Kumar1/go-bookstore-mngmnt/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Author      string `json:"author" gorm:"not null"`
	Publication string `json:"publication" gorm:"not null"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&Book{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

func (b *Book) CreateBook() *Book {
	if err := db.Create(&b).Error; err != nil {
		panic("Failed to create book: " + err.Error())
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	if err := db.Find(&Books).Error; err != nil {
		panic("Failed to fetch books: " + err.Error())
	}
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	if err := db.Where("ID=?", ID).Delete(&book).Error; err != nil {
		panic("Failed to delete book: " + err.Error())
	}
	return book
}
