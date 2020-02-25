package core

import (
	"fmt"
	"offers_iiko/config"
	"offers_iiko/core/server"
	"offers_iiko/service"
)

// Старт работы ядра
func Start(mode int) {
	// Загружаем конфигурацию
	setting := config.Load(mode)

	// проверяем  подклчение к бд
	service.SetConfig(setting)
	service.CheckStructure()

	// Запусткаем загрузку кофигураций политк безопастрости
	//access.PapEntity.Load()
	// обновляем  переменные для контроля доступа
	//access.PapEntity.Update()

	server.SetConfig(setting)
	server.Run(mode)

}

// запускаем сервер
func startServer() {
	fmt.Println("Запуск Сервера ")
}
