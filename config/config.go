package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/go-ini/ini"
)

const (
	DebugMode   = 0
	ReleaseMode = 1
)

type Config struct {
	Core          Core          `ini_section:"core" json:"core"`
	Client        SyncClient    `ini_section:"client" json:"client"`
	FileServer    FileServer    `ini_section:"file_server" json:"file_server"`
	ConfigWebhook ConfigWebhook `ini_section:"config_webhook" json:"config_webhook"`
}

type Core struct {
	DbDriver     string `ini:"db_driver" json:"db_driver"`
	DbSourceName string `ini:"db_source_name" json:"db_source_name"`
	ServerPort   string `ini:"server_port" json:"server_port"`
}
type SyncClient struct {
	Run                    bool   `ini:"run" json:"run"`
	CalculateOfferInClient bool   `ini:"calculate_offers_in_client" json:"calculate_offer_in_client"` //всегда считать акции на этой машине
	Interval               string `ini:"interval" json:"interval"`
	Token                  string `ini:"token" json:"token"`
	IP                     string `ini:"ip" json:"ip"`
	Port                   string `ini:"port" json:"port"`
	Debug                  string `ini:"debug" json:"debug"`
}
type FileServer struct {
	Path            string `ini:"path" json:"path"`
	ImageFolderName string `ini:"imagefoldername" json:"image_folder_name"`
	FilesFolderName string `ini:"filesfoldername" json:"files_folder_name"`
	Images          string `ini:"images" json:"images"`
}
type ConfigWebhook struct {
	Send            string `ini:"send" json:"send"`
	Interval        string `ini:"interval" json:"interval"`
	SucsessResponse string `ini:"sucsessResponse" json:"sucsess_response"`
}

var file = "setting.json"

// Загружает  настройки из файла core.setting.ini
// если файла нет , то создает новый с настройками по умолчанию
// если неудается открыть файл - паника
func Load(mode int) *Config {
	var path string
	if mode == ReleaseMode {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		path = dir + "/" + file

	} else {
		path = file
	}
	if !checkExist(path) {
		_, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		CreateJsonFile(path)
	}

	result := LoadJson(path)

	return result
}

func loadIni(path string) *Config {
	result := DefaultConfig
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	core := Core{}
	client := SyncClient{}
	fileServer := FileServer{}
	config_webhook := ConfigWebhook{}
	loadIniSection(cfg, "Core", &core)
	loadIniSection(cfg, "SyncClient", &client)
	loadIniSection(cfg, "FileServer", &fileServer)
	loadIniSection(cfg, "Config_webhook", &config_webhook)

	result.Core = core
	result.Client = client
	result.FileServer = fileServer
	result.ConfigWebhook = config_webhook
	return &result
}
func loadIniSectionReflect(cfg *ini.File, section string, entity reflect.Value) {

	sec := cfg.Section(section)
	typ := entity.Type()
	values := entity
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag
		tag_ini, is := tag.Lookup("ini")
		sec_value := sec.Key(tag_ini)
		if is {
			field := values.Field(i)
			switch field.Kind() {
			case reflect.Bool:
				{
					val, _ := sec_value.Bool()
					field.SetBool(val)
				}
			case reflect.Int:
				{
					val, _ := sec_value.Int()
					field.SetInt(int64(val))
				}
			case reflect.String:
				{
					field.SetString(sec_value.String())
				}
			}
		}

	}
}
func loadIniSection(cfg *ini.File, section string, entity interface{}) {

	sec := cfg.Section(section)
	typ := reflect.TypeOf(entity).Elem()
	values := reflect.ValueOf(entity).Elem()
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag
		tag_ini, is := tag.Lookup("ini")
		sec_value := sec.Key(tag_ini)
		if is {
			field := values.Field(i)
			switch field.Kind() {
			case reflect.Bool:
				{
					val, _ := sec_value.Bool()
					field.SetBool(val)
				}
			case reflect.Int:
				{
					val, _ := sec_value.Int()
					field.SetInt(int64(val))
				}
			case reflect.String:
				{
					field.SetString(sec_value.String())
				}
			}
		}

	}
}

func createDefault(path string) {
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	createDefaulSection(cfg, "Core", defaultCore)
	createDefaulSection(cfg, "SyncClient", defaultSyncClient)
	cfg.SaveTo(path)

}
func createDefaulSection(cfg *ini.File, section string, entity interface{}) {
	sec := cfg.Section(section)
	typ := reflect.TypeOf(entity)
	values := reflect.ValueOf(entity)

	for i := 0; i < typ.NumField(); i++ {

		tag := typ.Field(i).Tag
		tag_ini, is := tag.Lookup("ini")
		if is {
			field := values.Field(i)
			switch field.Kind() {
			case reflect.Bool:
				{
					value := strconv.FormatBool(field.Bool())
					fmt.Println(tag_ini, "write value:", value)
					sec.Key(tag_ini).SetValue(value)
				}
			case reflect.Int:
				{
					value := strconv.Itoa(int(field.Int()))
					sec.Key(tag_ini).SetValue(value)
				}
			case reflect.String:
				{
					sec.Key(tag_ini).SetValue(field.String())
				}
			}
			value := values.Field(i).String()
			sec.Key(tag_ini).SetValue(value)
		}

	}
}
func checkExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}
