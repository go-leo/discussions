package nacosx

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Configs map[string]*Config

type Config struct {
	Address   string `mapstructure:"address" json:"address" yaml:"address"`
	Port      uint64 `mapstructure:"port" json:"port" yaml:"port"`
	DataID    string `mapstructure:"data_id" json:"data_id" yaml:"data_id"`
	Group     string `mapstructure:"group" json:"group" yaml:"group"`
	Namespace string `mapstructure:"namespace" json:"namespace" yaml:"namespace"`
}

func NewConfigClients(configs Configs) (map[string]config_client.IConfigClient, error) {
	clients := make(map[string]config_client.IConfigClient)
	for key, config := range configs {
		client, err := NewConfigClient(config)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func NewRegistryClients(configs Configs) (map[string]naming_client.INamingClient, error) {
	clients := make(map[string]naming_client.INamingClient)
	for key, config := range configs {
		client, err := NewRegistryClient(config)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func NewConfigClient(config *Config) (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{*constant.NewServerConfig(config.Address, config.Port)}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(config.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("warn"),
	)
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
}

func NewRegistryClient(config *Config) (naming_client.INamingClient, error) {
	cc := constant.NewClientConfig(
		constant.WithNamespaceId(config.Namespace),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithUpdateCacheWhenEmpty(true),
	)
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.Address, config.Port),
	}
	clientParam := vo.NacosClientParam{ClientConfig: cc, ServerConfigs: sc}
	return clients.NewNamingClient(clientParam)
}
