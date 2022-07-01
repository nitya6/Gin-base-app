package utils

import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/zap"
)

type logConfig struct {
	Level            string            `json:"level"`
	Encoding         string            `json:"encoding"`
	OutputPaths      []string          `json:"outputPaths"`
	ErrorOutputPaths []string          `json:"errorOutputPaths"`
	EncoderConfig    map[string]string `json:"encoderConfig"`
}

var Logger *zap.Logger

// Initializes Logger with log configs
func InitializeLogger() {

	config := logConfig{
		Level:            os.Getenv("LOG_LEVEL"),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: map[string]string{
			"messageKey":   "message",
			"timeKey":      "T",
			"levelKey":     "level",
			"levelEncoder": "lowercase",
			"timeEncoder":  "RFC3339TimeEncoder",
		},
	}

	marshedJson, marshErr := json.Marshal(config)
	if marshErr != nil {
		log.Println("failed to configure logger", marshErr)
	}

	var cfg zap.Config
	if err := json.Unmarshal(marshedJson, &cfg); err != nil {
		log.Println("failed to configure logger", err)
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Println("failed to configure logger", err)
	}
	defer logger.Sync()

	logger.Info("logger configuration succeeded")
	Logger = logger
}
