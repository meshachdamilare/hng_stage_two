package database

import (
	"github/meshachdamilare/hng_stage_two/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Person{})
}

func CreatePerson(person *models.Person) error {
	if err := db.Create(person).Error; err != nil {
		return err
	}
	return nil
}

func GetPerson(id int) (*models.Person, error) {
	var person models.Person
	if err := db.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func UpdatePerson(person models.Person) models.Person {
	db.Save(&person)
	return person
}

func DeletePerson(id int) {
	db.Delete(&models.Person{}, "id = ?", id)
}
