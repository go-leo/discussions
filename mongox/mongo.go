package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Configs map[string]*Config

type Config struct {
	URI string `mapstructure:"uri" json:"uri" yaml:"uri"`
}

func NewClients(ctx context.Context, configs Configs) (map[string]*mongo.Client, error) {
	clients := make(map[string]*mongo.Client)
	for key, config := range configs {
		client, err := NewClient(ctx, config)
		if err != nil {
			return nil, err
		}
		clients[key] = client
	}
	return clients, nil
}

func NewClient(ctx context.Context, config *Config) (*mongo.Client, error) {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	return mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
}
