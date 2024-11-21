package scripts

import (
	"go-kas/config"
	"go-kas/services"
)

var cfg *config.Config

var service *services.Service

type Script struct {
	cfg *config.Config
}

func Migrate() {
	service = services.NewService(cfg)
}
