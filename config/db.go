package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/graphql-product/domain"
	"log"
	"log/slog"
	"os"
)

const dsnDefault = "root:korie123@tcp(127.0.0.1:3306)/graphql_product?charset=utf8mb4&parseTime=True&loc=Local"

func ConnDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = dsnDefault
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Errorf("error connect to mysql %v", err)
	}

	///

	if err = db.AutoMigrate(&domain.Category{}, &domain.Product{}); err != nil {
		fmt.Errorf("failed to auto migrate %v", err)
		panic(err)
	}
	slog.Info("success migrate ")
	return db
}
