package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser(name string) error {
	config.CurrentUserName = name
	return write(*config);
}

func Read () (*Config, error)  {
	config := &Config{}

	file, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	reader, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(reader, config); err != nil {
		return nil, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", homedir, configFileName), nil
}

func write(cfg Config) error {
	file, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return err
	}

	return nil
}