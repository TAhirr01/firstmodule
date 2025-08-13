package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := "host=localhost user=root password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Baku"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db baglantisi alinmadi: ", err)
	}
	DB = db

}
