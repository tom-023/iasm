package main

import (
	"time"

	"github.com/tom-023/iasm/config"
	"github.com/tom-023/iasm/logger"
	"go.uber.org/zap"
)

// StartTickerLoop starts a loop that runs the given task at specified intervals
func startTickerLoop(task func()) {
	c := config.GetConfig()
	intervalStr := c.GetString("monitor_interval")
	if intervalStr == "" {
		intervalStr = "5m"
		logger.Logger.Info("No monitor interval set, using default 5 minutes")
	}

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		logger.Logger.Fatal("Invalid monitor interval format", zap.String("interval", intervalStr), zap.Error(err))
		return
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	logger.Logger.Info("Starting task loop", zap.Duration("interval", interval))

	for {
		select {
		case <-ticker.C:
			logger.Logger.Info("Executing task")
			task()
		}
	}
}
