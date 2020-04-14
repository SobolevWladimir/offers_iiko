package config

var DefaultConfig = Config{
	Core:   defaultCore,
	Client: defaultSyncClient,
}
var defaultCore = Core{
	DbDriver:     "mysql",
	DbSourceName: "admin:111@tcp(10.8.0.3:3306)/altegra",
	ServerPort:   ":8091",
}
var defaultSyncClient = SyncClient{
	Run:      false,
	Interval: "3000",
	Token:    "",
	IP:       "127.0.0.1",
	Port:     "8080",
	Debug:    "0",
}
