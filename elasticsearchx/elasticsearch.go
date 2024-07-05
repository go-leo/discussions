package elasticsearchx

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Configs map[string]*Config

type Config struct {
	Addresses []string `mapstructure:"addresses" json:"addresses" yaml:"addresses"`
	Username  string   `mapstructure:"username" json:"username" yaml:"username"`
	Password  string   `mapstructure:"password" json:"password" yaml:"password"`
}

func NewClients(configs Configs) (map[string]*elasticsearch.Client, error) {
	clients := make(map[string]*elasticsearch.Client)
	for key, config := range configs {
		client, err := NewClient(config)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func NewClient(config *Config) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Addresses,
		Username:  config.Username,
		Password:  config.Password,
	})
}
