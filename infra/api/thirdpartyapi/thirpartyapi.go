package thirdpartyapi

import (
	"fmt"
	"net/http"
)

type Config struct {
	BaseURL   string `env:"BASEURL,default=http://localhost"`
	ClientID  string `env:"CLIENT_ID,default=clientId"`
	SecretKey string `env:"SECRET_KEY,default=secretkey"`
	Port      int    `env:"PORT,default=8000"`
}

type Client struct {
	Config    *Config `env:"BASEURL"`
	ClientID  string  `env:"CLIENT_ID"`
	SecretKey string  `env:"SECRET_KEY"`
}

func NewClient(config *Config) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) GetBaseURL() string {
	if c.Config.Port > 0 {
		return fmt.Sprintf("%s%s", c.Config.BaseURL, fmt.Sprintf(":%d", c.Config.Port))
	}
	return c.Config.BaseURL
}

func (c *Client) Get() (int, error) {
	response, err := http.Get(fmt.Sprintf("%s/health", c.GetBaseURL()))

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}
