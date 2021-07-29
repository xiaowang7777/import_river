package config

import (
	"gopkg.in/yaml.v3"
	"import_river/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

const Config_File_Name = "river_config.yaml"

type Config struct {
	FilePath     string `json:"file_path" yaml:"file_path,omitempty"`
	Host         string `json:"host" yaml:"host,omitempty"`
	Port         int16  `json:"port" yaml:"port,omitempty"`
	Username     string `json:"username" yaml:"username,omitempty"`
	Password     string `json:"password" yaml:"password,omitempty"`
	DatabaseName string `json:"database_name" yaml:"database_name,omitempty"`
	TableName    string `json:"table_name" yaml:"table_name,omitempty"`
}

func New() (*Config, error) {
	c := &Config{}
	err := c.Load()
	return c, err
}

func (c *Config) Load() error {
	err := createIfNotExists()
	if err != nil {
		return err
	}
	file, err := os.Open(getConfigFilePath())
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if len(bytes) > 0 {
		return yaml.Unmarshal(bytes, c)
	}
	return nil
}

func (c *Config) Write() error {
	file, err := os.OpenFile(getConfigFilePath(), os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	return encoder.Encode(c)
}

func createIfNotExists() error {
	filePath := filepath.Join(utils.HomeDir(), Config_File_Name)
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		_, err := os.Create(filePath)
		return err
	} else if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() string {
	filePath := filepath.Join(utils.HomeDir(), Config_File_Name)
	if _, err := os.Stat(filePath); err == nil {
		return filePath
	}
	return ""
}
