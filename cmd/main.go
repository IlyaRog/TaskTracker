package main

import (
	internal "TaskTracker/internal"
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Config struct {
	Path string `json:"path"`
}

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "Path to config file")
	flag.Parse()

	config := loadConfig(configPath)

	if err := os.MkdirAll(config.Path, os.ModePerm); err != nil {
		log.Fatalf("не удалось создать каталог данных: %v", err)
	}

	app := internal.CreateCLI(&internal.JsonTaskRepository{FilePath: config.Path})
	app.Menu()

}

func loadConfig(path string) *Config {
	fileData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("не удалось прочитать конфиг %q: %v", path, err)
	}

	config := &Config{}
	if err := json.Unmarshal(fileData, config); err != nil {
		log.Fatalf("ошибка парсинга конфигурации: %v", err)
	}

	return config
}
