package database

import (
	"be17/cleanarch/app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection
func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	// declare struct config & variable connectionString
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitialMigration() {
	// DB.AutoMigrate(&User{})
	// auto migrate untuk table book
}
