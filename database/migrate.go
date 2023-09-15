package database

import (
	"github.com/shiv122/go-test/app/models"
	"github.com/shiv122/go-test/config"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{})
}
