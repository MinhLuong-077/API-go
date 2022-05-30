package setup

import (
	"fmt"

	"gorm.io/gorm"
	"stockexchange.com/entity"
)

func Init(db *gorm.DB) {



	if !db.Migrator().HasTable(&entity.Transaction{}) {
		db.AutoMigrate(&entity.Transaction{})
	}
	if !db.Migrator().HasTable(&entity.Bankaccount{}) {
		db.AutoMigrate(&entity.Bankaccount{})
	}

	fmt.Println("Database Initialized!")
}

