package mysql

import (
	"fmt"
	"log"
	"technical-test/internal/domain/entity"
	"technical-test/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// GetDBConnection gets DB connection
func NewMysqlConnection() *gorm.DB {
	configEnv := config.LoadEnv()
	loadConfig := map[string]string{
		"Username": configEnv.GetString("DB_USERNAME"),
		"Passowrd": configEnv.GetString("DB_PASSWORD"),
		"Host":     configEnv.GetString("DB_HOST"),
		"Port":     configEnv.GetString("DB_PORT"),
		"DB":       configEnv.GetString("DB_NAME"),
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", loadConfig["Username"], loadConfig["Password"], loadConfig["Host"], loadConfig["Port"], loadConfig["DB"])
	log.Println("Initialize DB connection...")
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&entity.Class{},
		&entity.Student{},
	)
	return db
}
