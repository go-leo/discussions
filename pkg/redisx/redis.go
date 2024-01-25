package redisx

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"time"
)

type Configs map[string]*Config

// Config 配置信息
type Config struct {
	// Either a single address or a seed list of host:port addresses
	// of cluster/sentinel nodes.
	Addrs []string `mapstructure:"addrs" json:"addrs" yaml:"addrs"`

	// Only cluster clients.
	ClusterClientsOptions ClusterClientsOptions

	// Only failover clients.
	FailOverClients FailOverClients

	// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
	ClientName string `mapstructure:"client_name" json:"client_name" yaml:"client_name"`

	// Database to be selected after connecting to the server.
	// Only single-node and failover clients.
	DB int `mapstructure:"db" json:"db" yaml:"db"`

	// Common options.
	Protocol         int    `mapstructure:"protocol" json:"protocol" yaml:"protocol"`
	Username         string `mapstructure:"username" json:"username" yaml:"username"`
	Password         string `mapstructure:"password" json:"password" yaml:"password"`
	SentinelUsername string `mapstructure:"sentinel_username" json:"sentinel_username" yaml:"sentinel_username"`
	SentinelPassword string `mapstructure:"sentinel_password" json:"sentinel_password" yaml:"sentinel_password"`

	MaxRetries      int           `mapstructure:"max_retries" json:"max_retries" yaml:"max_retries"`
	MinRetryBackoff time.Duration `mapstructure:"min_retry_backoff" json:"min_retry_backoff" yaml:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `mapstructure:"max_retry_backoff" json:"max_retry_backoff" yaml:"max_retry_backoff"`

	DialTimeout           time.Duration `mapstructure:"dial_timeout" json:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout           time.Duration `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout          time.Duration `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
	ContextTimeoutEnabled bool          `mapstructure:"context_timeout_enabled" json:"context_timeout_enabled" yaml:"context_timeout_enabled"`

	// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
	PoolFIFO bool `mapstructure:"pool_fifo" json:"pool_fifo" yaml:"pool_fifo"`

	PoolSize        int           `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size"`
	PoolTimeout     time.Duration `mapstructure:"pool_timeout" json:"pool_timeout" yaml:"pool_timeout"`
	MinIdleConns    int           `mapstructure:"min_idle_conns" json:"min_idle_conns" yaml:"min_idle_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time" yaml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime" yaml:"conn_max_lifetime"`

	// otel
	EnableTracing bool `mapstructure:"enable_tracing" json:"enable_tracing" yaml:"enable_tracing"`
	EnableMetrics bool `mapstructure:"enable_metrics" json:"enable_metrics" yaml:"enable_metrics"`

	TLSConfig *tls.Config
}

type FailOverClients struct {
	// MasterName The sentinel master name.
	MasterName string `mapstructure:"master_name" json:"master_name" yaml:"master_name"`
}

type ClusterClientsOptions struct {
	MaxRedirects   int  `mapstructure:"max_redirects" json:"max_redirects" yaml:"max_redirects"`
	ReadOnly       bool `mapstructure:"read_only" json:"read_only" yaml:"read_only"`
	RouteByLatency bool `mapstructure:"route_by_latency" json:"route_by_latency" yaml:"route_by_latency"`
	RouteRandomly  bool `mapstructure:"route_randomly" json:"route_randomly" yaml:"route_randomly"`
}

// New 创建redis客户端集合
func NewRedisClients(ctx context.Context, configs Configs) (map[string]redis.UniversalClient, error) {
	clients := make(map[string]redis.UniversalClient)
	for key, o := range configs {
		client := redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:                 o.Addrs,
			ClientName:            o.ClientName,
			DB:                    o.DB,
			Protocol:              o.Protocol,
			Username:              o.Username,
			Password:              o.Password,
			SentinelUsername:      o.SentinelUsername,
			SentinelPassword:      o.SentinelPassword,
			MaxRetries:            o.MaxRetries,
			MinRetryBackoff:       o.MinRetryBackoff,
			MaxRetryBackoff:       o.MaxRetryBackoff,
			DialTimeout:           o.DialTimeout,
			ReadTimeout:           o.ReadTimeout,
			WriteTimeout:          o.WriteTimeout,
			ContextTimeoutEnabled: o.ContextTimeoutEnabled,
			PoolFIFO:              o.PoolFIFO,
			PoolSize:              o.PoolSize,
			PoolTimeout:           o.PoolTimeout,
			MinIdleConns:          o.MinIdleConns,
			MaxIdleConns:          o.MaxIdleConns,
			ConnMaxIdleTime:       o.ConnMaxIdleTime,
			ConnMaxLifetime:       o.ConnMaxLifetime,
			TLSConfig:             o.TLSConfig,
			MaxRedirects:          o.ClusterClientsOptions.MaxRedirects,
			ReadOnly:              o.ClusterClientsOptions.ReadOnly,
			RouteByLatency:        o.ClusterClientsOptions.RouteByLatency,
			RouteRandomly:         o.ClusterClientsOptions.RouteRandomly,
			MasterName:            o.FailOverClients.MasterName,
		})
		_, err := client.Ping(ctx).Result()
		if err != nil {
			return nil, fmt.Errorf("failed ping redis %v, %w", key, err)
		}
		// Enable tracing instrumentation.
		if o.EnableTracing {
			if err := redisotel.InstrumentTracing(client); err != nil {
				return nil, err
			}
		}
		if o.EnableMetrics {
			if err := redisotel.InstrumentMetrics(client); err != nil {
				return nil, err
			}
		}
		clients[key] = client
	}
	return clients, nil
}
