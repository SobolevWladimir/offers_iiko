package config

var DefaultConfig = Config{
	Core:          defaultCore,
	Client:        defaultSyncClient,
	FileServer:    defaultFileServer,
	ConfigWebhook: defaultConfig_webhook,
}
var defaultCore = Core{
	DbDriver:     "mysql",
	DbSourceName: "admin:111@tcp(192.168.1.48:3306)/altegra",
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
var defaultConfig_webhook = ConfigWebhook{
	Send:            "3",
	Interval:        "1",
	SucsessResponse: "200,201",
}
var defaultFileServer = FileServer{
	Path:            "./files",
	ImageFolderName: "images",
	FilesFolderName: "files",
	Images:          "png,jpeg,jpg",
}
