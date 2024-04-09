package utils

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatalf("Failed to open log file: %v", err)
	}

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05.000",
		DisableLevelTruncation: true,
		DisableColors:          true,
	})

	Logger.SetOutput(file)

	err = godotenv.Load()
	if err != nil {
		Logger.Errorf("Error loading .env file: %v", err)
	} else {
		Logger.Info(".env file loaded successfully")
	}

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		Logger.Fatalf("Error parsing log level: %v", err)
	} else {
		Logger.Infof("Log level set to: %s", logLevel.String())
	}

	Logger.SetLevel(logLevel)
}
