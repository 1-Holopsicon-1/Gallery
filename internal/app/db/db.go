package db

import (
	"Gallery/internal/app/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {
	e := godotenv.Load()
	if e != nil {
		panic("No env")
	}

	user := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := os.Getenv("db_port")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		panic("Error connection")
	}
	return db

}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
		panic("Cant migrate User")
	}

	err = db.AutoMigrate(&models.Tags{})
	if err != nil {
		fmt.Println(err)
		panic("Cant migrate Tags")
	}

	err = db.AutoMigrate(&models.Posts{})
	if err != nil {
		fmt.Println(err)
		panic("Cant migrate Post")

	}
}
