package config

import (
	"errors"
	"os"

	"github.com/Ja7ad/fluora/pkg/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Rest   *Rest          `yaml:"rest"`
	AI     []*AI          `yaml:"ai"`
	Logger *logger.Config `yaml:"logger"`
}

type Rest struct {
	Address string
	Origins []string `yaml:"origins"`
}

type AI struct {
	Provider  string `yaml:"provider"`
	Serve     string `yaml:"serve"`
	APIKey    string `yaml:"api_key"`
	RateLimit int    `yaml:"rate_limit"`
}

func LoadConfig(cfgPath string) (*Config, error) {
	f, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(f, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) BasicCheck() error {
	if c.Logger == nil {
		c.Logger = logger.DefaultConfig()
	}

	if len(c.AI) == 0 {
		return errors.New("no AI configuration found")
	}

	if c.Rest == nil {
		return errors.New("no REST configuration found")
	}

	if c.Rest.Address == "" {
		return errors.New("no REST configuration found")
	}

	return nil
}
