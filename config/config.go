package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	//embeded struct
	Server struct {
		Port string `ymal:"port"`
	} `yaml:"server"`
}

var AppConfig Config

func GetConfig() error {

	//read file by os
	file, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}
