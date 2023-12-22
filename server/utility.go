package main

import (
	"log"

	"go.uber.org/zap"
)

const (
	developmentModeLoggingMask = 0x01
)

func isDevelopmentModeLogging(featureFlags uint32) bool {
	return (featureFlags & developmentModeLoggingMask) != 0
}

func zapSetup(developmentMode bool) *zap.SugaredLogger {
	var err error
	var logger *zap.Logger

	atomic := zap.NewAtomicLevel()
	atomic.SetLevel(zap.DebugLevel)

	if developmentMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	return logger.Sugar()
}
