package daabase

import (
	"log"
	"os"

	"bitbucket.org/nithursika/go-fiber-tutorial/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	//addmigration
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DBInstance{Db: db}
}
