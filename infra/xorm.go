package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"king-tool-api/model"
)

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Users{})
	return db
}

func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://postgres:password@db:5432/king-tool-api?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return db
}
