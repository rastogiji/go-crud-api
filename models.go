package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Art struct {
	gorm.Model
	ArtName       string `json:"artname"`
	Description   string `json:"description"`
	CoverImageUrl string `json:"coverimageurl"`
}

var DB *gorm.DB

func ConnectToDB() {

	DBUser := os.Getenv("DBUSER")
	DBPassword := os.Getenv("DBPASSWORD")
	DBHost := os.Getenv("DBHOST")
	DBName := os.Getenv("DBNAME")
	DBUrl := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", DBUser, DBPassword, DBHost, DBName)
	d, err := gorm.Open(mysql.Open(DBUrl), &gorm.Config{})

	if err != nil {
		panic("Unable to Connect to DB")
	}
	fmt.Println("Connected to DB")
	DB = d

	if err := DB.AutoMigrate(&Art{}); err != nil {
		log.Fatal("Unable to Automigrate Schema")
	}
}
