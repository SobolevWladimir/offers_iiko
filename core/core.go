package core

import (
	"altegra_offers/config"
	"altegra_offers/core/server"
	"altegra_offers/service"
	"fmt"
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
