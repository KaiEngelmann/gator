package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	gatorconfig := home_dir + "/.gatorconfig.json"
	return gatorconfig, nil
}

func Read() (*Config, error) {
	gatorconfig, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(gatorconfig)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	gatorconfig, _ := getConfigFilePath()
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	permissions := os.FileMode(0644)
	err = os.WriteFile(gatorconfig, data, permissions)
	if err != nil {
		return err
	}
	return nil
}
