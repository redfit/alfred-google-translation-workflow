package main

import (
	"os"
	"path/filepath"
	"encoding/json"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type config struct {
	AccessToken string `json:"accessToken"`
}

func getDefaultConfigPath() string {
	homeDir, _ := homedir.Dir()
	return filepath.Join(homeDir, "Library/Application Support/Alfred 2/Workflow Data/", bundleID)
}

func loadConfig() error {
	configPath := getDefaultConfigPath()
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func saveConfig() error {

	var marshaledConfig config

	configPath := getDefaultConfigPath()
	viper.Unmarshal(&marshaledConfig)

	buf, err := json.MarshalIndent(marshaledConfig, "", "    ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(configPath, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(configPath, "config.json"))
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(string(buf))
	return nil
}
