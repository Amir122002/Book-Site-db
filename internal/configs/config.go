package configs

import (
	"book/pkg/models"
	"encoding/json"
	"log"
	"os"
)

func InitConfig() (*models.Config, error) {

	bytes, err := os.ReadFile("./internal/configs/config.json")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var config models.Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &config, nil
}
