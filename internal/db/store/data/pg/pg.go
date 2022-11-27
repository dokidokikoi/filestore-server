package postgres

import (
	"fmt"

	"file-store/internal/config"
	"file-store/internal/db/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store struct {
	db *gorm.DB
}

func GetPGFactory(pgconf *config.PGConfig) (*Store, error) {
	var err error
	dns := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=%s password=%s", pgconf.Host, pgconf.UserName, pgconf.Database, pgconf.Port, pgconf.TimeZone, pgconf.Password)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get postgresql store factory")
	}

	migrateDatabase(db)

	return &Store{db}, nil
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&model.File{})
}
