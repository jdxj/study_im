package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func PrintConfigFormat() ([]byte, error) {
	data, _ := yaml.Marshal(Config{})

	fileName := "conf.yaml.default"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	defer file.Sync()

	_, err = file.Write(data)
	return data, err
}

func New(path string) (*Config, error) {
	conf := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return conf, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	return conf, decoder.Decode(conf)
}

type Config struct {
	Logger Logger `yaml:"logger"`
	Gate   Gate   `yaml:"gate"`
}

type Logger struct {
	FileName   string `yaml:"file_name"`
	AppName    string `yaml:"app_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Level      int    `yaml:"level"`
	LocalTime  bool   `yaml:"local_time"`
	Compress   bool   `yaml:"compress"`
}

type Gate struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
