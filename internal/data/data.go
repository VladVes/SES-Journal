package data

import (
	"log"
	"time"

	"github.com/sirupsen/logrus"
	// conf "github.com/VladVes/SES-Journal/internal/config"
	appLogger "github.com/VladVes/SES-Journal/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	DBLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	MaxOpenConns    = 10
	ConnMaxLifetime = time.Hour
)

func makeLogger() *DBLogger.Interface {
	newLogger := DBLogger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		DBLogger.Config{
			SlowThreshold: time.Second,
			LogLevel:      DBLogger.Info,
			Colorful:      true,
		},
	)
	return &newLogger
}

func getDBConfig() *gorm.Config {
	dbLogger := makeLogger()
	return &gorm.Config{
		Logger:                 *dbLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "app_",
		},
	}
}

func DBConnet(dsn string) (*gorm.DB, error) {
	dbConf := getDBConfig()
	db, err := gorm.Open(postgres.Open(dsn), dbConf)
	if err != nil {
		// log.Fatalf("DB connection error: %v", err)
		return nil, err
	}

	appLogger.Log.WithFields(logrus.Fields{
		"dsn": dsn,
	}).Info("DB connection established")

	sqlDB, err := db.DB()
	if err != nil {
		// log.Fatalf("DB pool error, %v", err)
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		// log.Fatalf("DB ping error: %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(MaxOpenConns)
	sqlDB.SetConnMaxLifetime(ConnMaxLifetime)

	appLogger.Log.WithFields(logrus.Fields{
		"MaxOpenConns":    MaxOpenConns,
		"ConnMaxLifeTime": ConnMaxLifetime,
	}).Info("DB connection pool configured")

	return db, nil
}
