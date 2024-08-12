package config

import (
	"io"
	"os"

	"github.com/pelletier/go-toml"
)

var Port int
var Host string
var DebugMode bool
var JWT JWTConfig
var DB DatabaseConfig

func LoadConfig() {
	file, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config Config
	err = toml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	Port = config.Port
	Host = config.Host
	DebugMode = config.DebugMode
	JWT = config.JWT
	DB = config.Database
}
