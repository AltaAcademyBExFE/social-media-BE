package database

import (
	"fmt"
	"sosmed/config"
	rc "sosmed/features/comment/repository"
	rp "sosmed/features/post/repository"
	ur "sosmed/features/user/repository"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPwd,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&rp.Post{})
	db.AutoMigrate(&rc.Comment{})
	db.AutoMigrate(&ur.User{})
}
