package initializers

import (
	"fmt"
	"goselflearn/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DBConnect() {
	var err error

	sslMode := "disable"
	if Config.DBSSLMode {
		sslMode = "enable"
	}
	postgresDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tehran",
		Config.DBHost,
		Config.DBUser,
		Config.DBPass,
		Config.DBName,
		Config.DBPort,
		sslMode,
	)

	DB, err = gorm.Open(
		postgres.Open(postgresDSN),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Info),
			TranslateError: true,
		},
	)
	if err != nil {
		log.Fatalln("failed to connect to the database, dsn:", postgresDSN, " error:", err.Error())
	}

	err = DB.AutoMigrate(
		models.User{},
		models.Post{},
	)
	if err != nil {
		log.Fatalln("failed to migrate models. error:", err.Error())
	}
}
