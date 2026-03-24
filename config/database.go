package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	uriDB := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		MainConfig.DBConnection,
		MainConfig.DBUsername,
		MainConfig.DBPassword,
		MainConfig.DBHost,
		MainConfig.DBPort,
		MainConfig.DBName,
	)

	fmt.Println(uriDB)

	db, err := gorm.Open(postgres.Open(uriDB), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(MainConfig.DBMaxIdleConnection)
	sqlDB.SetMaxOpenConns(MainConfig.DBMaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(MainConfig.DBMaxLifetimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(MainConfig.DBMaxIdleTime) * time.Second)

	// Auto migrate
	err = db.AutoMigrate(
	// &models.User{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
