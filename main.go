package main

import (
	"github.com/viitorags/encurtadorUrl/config"
	"github.com/viitorags/encurtadorUrl/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	if err := config.InitConfig(); err != nil {
		logger.Error("Erro ao inicializar configuração:", err)
		return
	}

	router.InitRoutes()

}
