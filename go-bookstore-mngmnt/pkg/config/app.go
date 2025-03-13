package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// First connect without database to create it if it doesn't exist
	dsn := "root:kumarshubh16@tcp(localhost:3306)/"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MySQL: %v", err))
	}

	// Create database if it doesn't exist
	sqlDB, err := d.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get database instance: %v", err))
	}
	
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS bookstore")
	if err != nil {
		panic(fmt.Sprintf("Failed to create database: %v", err))
	}

	// Now connect to the specific database
	dsn = "root:kumarshubh16@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
