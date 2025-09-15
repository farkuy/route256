package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert/yaml"
)

type Config struct {
	Token string `yaml: token`
}

type Service struct {
	config Config
}

func (c *Service) GetToken() string {
	return c.config.Token
}

// TODO:Вынести это ублюдство в env
const path = "C:/Users/Onton/route256/data/config.yaml"

func New() (*Service, error) {
	rows, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "open config file")
	}

	var config *Config
	err = yaml.Unmarshal(rows, &config)
	if err != nil {
		return nil, errors.Wrap(err, "parse config")
	}

	return &Service{config: *config}, nil
}
