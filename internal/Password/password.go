package password

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type DB struct {
	DB gorm.DB
}

type Password struct {
	ID           int       `json:"id" gorm:"column:id;type:int;autoIncrement;not null"`
	Name         string    `json:"name" gorm:"column:name;type:text;not null"`
	Value        string    `json:"value" gorm:"column:value;type:text;not null"`
	Category     string    `json:"category" gorm:"column:category;type:text;not null"`
	CreateAt     time.Time `json:"create_at" gorm:"column:create_at;type:timestamp;not null"`
	LastModified time.Time `json:"last_modified" gorm:"column:last_modified;type:timestamp;not null"`
}

func (db *DB) NewPassword(password Password) error {
	result := db.DB.Create(&password)
	if result.Error != nil {
		return fmt.Errorf("error insert password: %v", result.Error)
	}
	log.Println("Password save ✅")
	return nil
}

func (db *DB) ReadPassword(passwords *[]Password) error {
	res := db.DB.Find(&passwords)
	if res.Error != nil {
		return fmt.Errorf("failed reading passwords from DataBase: %v", res.Error)
	}

	return nil
}
