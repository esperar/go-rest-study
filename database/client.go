package database

import (
	"github.com/jinzhu/gorm"
	"gorest/entity"
	"log"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	log.Println("Connection was successful")
	return nil
}

//Migrate create/updates database table
func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
