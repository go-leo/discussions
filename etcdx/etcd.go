package etcdx

import (
	"go.etcd.io/etcd/client/v3"
)

type Configs map[string]*Config

type Config struct {
	Endpoints []string `mapstructure:"endpoints" json:"endpoints" yaml:"endpoints"`
	Username  string   `mapstructure:"username" json:"username" yaml:"username"`
	Password  string   `mapstructure:"password" json:"password" yaml:"password"`
}

func NewClients(configs Configs) (map[string]*clientv3.Client, error) {
	clients := make(map[string]*clientv3.Client)
	for key, config := range configs {
		client, err := NewClient(config)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func NewClient(config *Config) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: config.Endpoints,
		Username:  config.Username,
		Password:  config.Password,
	})
}
