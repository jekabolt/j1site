package jsite

import (
	"encoding/json"
	"os"
)

// Coifig is a configuration file
type Config struct {
	Name      string   `json:"name"`
	Host      string   `json:"host"`
	Port      string   `json:"port"`
	Filespath string   `json:"filespath"`
	Templates []string `json:"templates"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return Config{}, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
