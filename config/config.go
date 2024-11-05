package config

import (
	"github.com/spf13/viper"
	"strings"
)

var c *viper.Viper

// Init initializes config
func Init() {
	c = viper.New()
	c.AutomaticEnv()
	// 環境変数名とViperキーの形式を統一
	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// GetConfig returns config
func GetConfig() *viper.Viper {
	return c
}
