package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	if config.DBInitial == "POSTGREE" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	} else if config.DBInitial == "MYSQL" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	} else {
		log.Fatal("Failed to connect to the Database")
	}
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
