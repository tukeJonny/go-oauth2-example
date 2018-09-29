package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	Nonce    string     `yaml:"nonce"`
	AuthConf AuthConfig `yaml:"auth"`
}

type AuthConfig struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Scopes       []string `yaml:"scopes"`
	RedirectURL  string   `yaml:"redirect_url"`
}

func LoadConfig(path string) (*AppConfig, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var appConf AppConfig
	if err := yaml.Unmarshal(b, &appConf); err != nil {
		return nil, err
	}

	return &appConf, nil
}
