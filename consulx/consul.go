package consulx

import (
	"github.com/hashicorp/consul/api"
)

type Configs map[string]*Config

type Config struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"`
}

func NewClients(configs Configs) (map[string]*api.Client, error) {
	clients := make(map[string]*api.Client)
	for key, config := range configs {
		client, err := NewClient(config)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func NewClient(config *Config) (*api.Client, error) {
	return api.NewClient(&api.Config{
		Address: config.Address,
	})
}
