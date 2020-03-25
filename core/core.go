package core

import (
	"fmt"
	"offers_iiko/config"
	"offers_iiko/core/server"
	"offers_iiko/demon"
	"offers_iiko/service"
)

// Старт работы ядра
func Start(mode int) {
	// Загружаем конфигурацию
	setting := config.Load(mode)

	// проверяем  подклчение к бд
	service.SetConfig(setting)
	service.CheckStructure()

	demon.Start()

	server.SetConfig(setting)
	server.Run(mode)

}

// запускаем сервер
func startServer() {
	fmt.Println("Запуск Сервера ")
}
