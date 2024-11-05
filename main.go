package main

import (
	"github.com/tom-023/iasm/config"
	"github.com/tom-023/iasm/logger"
)

func main() {
	config.Init()
	logger.Init()
	defer logger.Sync()
}
