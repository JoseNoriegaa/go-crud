package main

// Author: Jose Noriega
// email: josenoriega723@gmail.com
// Last Update: 2020-02-16

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	cli "github.com/josenoriegaa/go-crud/cli"
	structs "github.com/josenoriegaa/go-crud/structs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cnnStr := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open("mysql", cnnStr)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&structs.Student{})

	option := -1
	for option != 6 {
		option = cli.Menu()
		switch option {
		case 1:
			cli.List(db)
			break
		case 2:
			cli.Find(db)
			break
		case 3:
			cli.CaptureStudent(db)
			break
		case 4:
			cli.Update(db)
		case 5:
			cli.Delete(db)
			break
		}
	}
}
