package commons

import (
	"apiGoMysql/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1098741116@/testapigo?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("migrando...")
	db.AutoMigrate(&models.Persona{})
}
