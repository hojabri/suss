package sqlite

import (
	"github.com/hojabri/suss/pkg/config"
	"github.com/hojabri/suss/pkg/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)
var DB *gorm.DB

func Connect() error {
	var err error

	logLevelConfig := config.Config.GetString("DB.GORM_LOG_LEVEL")
	var logLevel logger.LogLevel

	switch logLevelConfig {
	case "info":
		logLevel = logger.Info
	case "error":
		logLevel = logger.Error
	default:
		logLevel = logger.Error
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      true,
		},
	)
	DB, err = gorm.Open(sqlite.Open("db/suss.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err!=nil {
		return err
	}

	//Create tables
	err = DB.AutoMigrate(&entities.Event{})
	if err != nil {
		return err
	}
	return nil

}