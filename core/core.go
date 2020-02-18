package core

import (
	"fmt"
	"offers_iiko/config"
	"offers_iiko/core/server"
	"offers_iiko/service"
)

// Старт работы ядра
func Start() {
	// Загружаем конфигурацию
	setting := config.Load()

	// проверяем  подклчение к бд
	service.SetConfig(setting)
	service.CheckStructure()

	// Запусткаем загрузку кофигураций политк безопастрости
	//access.PapEntity.Load()
	// обновляем  переменные для контроля доступа
	//access.PapEntity.Update()

	server.SetConfig(setting)
	server.Run()

}

// запускаем сервер
func startServer() {
	fmt.Println("Запуск Сервера ")
}
