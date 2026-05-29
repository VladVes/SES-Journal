package repository

import (
	"log"
	"time"

	conf "github.com/VladVes/SES-Journal/internal/config"
	appLogger "github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/models"

	"github.com/sirupsen/logrus"
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

func CountEntities(db *gorm.DB, model any) int64 { //TODO: model any - so-so
	var counter int64
	if err := db.Model(model).Count(&counter).Error; err != nil {
		log.Fatalf("Seeds DB users count error: %v", err)
	}

	return counter
}

func seedDB(db *gorm.DB) {
	var user models.User
	var entry models.Entry

	usersCount := CountEntities(db, &user)
	entryCount := CountEntities(db, &entry)

	// TODO: logic dublication
	if usersCount == 0 {
		if err := db.Exec(UsersSeedsQuery).Error; err != nil {
			log.Fatalf("Seeds DB users error: %v", err)
		}
		appLogger.Log.WithFields(logrus.Fields{
			"usersCount": CountEntities(db, &user),
		}).Info("DB users seeds successful")
	}

	if entryCount == 0 {
		if err := db.Exec(EnriesSeedQuery).Error; err != nil {
			log.Fatalf("Entry Seeds error: %v", err)
		}
		appLogger.Log.WithFields(logrus.Fields{
			"entryCount": CountEntities(db, &entry),
		}).Info("DB entry seeds successful")
	}
}

func initDBDevMode(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Entry{}); err != nil {
		log.Fatalf("DB development mode initialization error: %v", err)
	}
	appLogger.Log.WithFields(logrus.Fields{
		"ENV": conf.DevEnv,
	}).Info("DB development mode initialization successful")

	seedDB(db)
}

func DBConnet(dsn, env string) (*gorm.DB, error) {
	dbConf := getDBConfig()
	db, err := gorm.Open(postgres.Open(dsn), dbConf)
	if err != nil {
		return nil, err
	}

	appLogger.Log.WithFields(logrus.Fields{
		"dsn": dsn,
	}).Info("DB connection established")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(MaxOpenConns)
	sqlDB.SetConnMaxLifetime(ConnMaxLifetime)

	appLogger.Log.WithFields(logrus.Fields{
		"MaxOpenConns":    MaxOpenConns,
		"ConnMaxLifeTime": ConnMaxLifetime,
	}).Info("DB connection pool configured")

	log.Println("DB connection established")

	if env == conf.DevEnv {
		initDBDevMode(db)
	}

	if env == conf.DevEnv {
		initDBDevMode(db)
	}

	return db, nil
}
