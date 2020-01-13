package main

import (
	"fmt"

	"github.com/waltton/confload"
)

// Config struct with params to start the application
type Config struct {
	Database struct {
		Name string `conf:"name" conf-usage:"database name"`
	} `conf:"database"`
	Server struct {
		Addr string `conf:"addr"`
		Port int    `conf:"port"`
	} `conf:"server"`
}

func main() {
	cfg := Config{}

	cfg.Database.Name = "123"
	cfg.Server.Addr = "456"
	cfg.Server.Port = 789

	confload.Load(&cfg, confload.FlagLoader)

	fmt.Printf("cfg: %+v", cfg)
}
