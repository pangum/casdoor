package core

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type Client struct {
	*casdoorsdk.Client
}

func NewClient(config *Config) *Client {
	return &Client{
		Client: casdoorsdk.NewClient(
			config.Endpoint,
			config.Id, config.Secret, config.Certificate,
			config.Organization, config.Application,
		),
	}
}
