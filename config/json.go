package config

import (
	"encoding/json"
	"io/ioutil"
)

func CreateJsonFile(path string) {
	bytes, err := json.MarshalIndent(&DefaultConfig, "", "\t")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(path, bytes, 0770)
}
func LoadJson(path string) *Config {
	result := new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		panic(err)
	}
	return result
}
