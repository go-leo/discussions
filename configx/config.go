package configx

import (
	"github.com/go-leo/pkg/actuatorx"
	"github.com/go-leo/pkg/amqpx"
	"github.com/go-leo/pkg/consulx"
	"github.com/go-leo/pkg/databasex"
	"github.com/go-leo/pkg/elasticsearchx"
	"github.com/go-leo/pkg/etcdx"
	"github.com/go-leo/pkg/grpcx"
	"github.com/go-leo/pkg/kafkax"
	"github.com/go-leo/pkg/mongox"
	"github.com/go-leo/pkg/nacosx"
	"github.com/go-leo/pkg/otelx"
	"github.com/go-leo/pkg/redisx"
)

var AppConf Configuration

type EnvConfig struct {
	Name      string `mapstructure:"SERVICE_NAME" json:"SERVICE_NAME" yaml:"SERVICE_NAME"`
	Namespace string `mapstructure:"SERVICE_NAMESPACE" json:"SERVICE_NAMESPACE" yaml:"SERVICE_NAMESPACE"`
	ID        string `mapstructure:"SERVICE_ID" json:"SERVICE_ID" yaml:"SERVICE_ID"`
	Version   string `mapstructure:"SERVICE_VERSION" json:"SERVICE_VERSION" yaml:"SERVICE_VERSION"`
}

type Configuration struct {
	Env           *EnvConfig             `mapstructure:"env" json:"env" yaml:"env"`
	Actuator      *actuatorx.Config      `mapstructure:"actuator" json:"actuator" yaml:"actuator"`
	Nacos         nacosx.Configs         `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	GrpcServer    *grpcx.ServerConfig    `mapstructure:"grpc_server" json:"grpc_server" yaml:"grpc_server"`
	GrpcClient    grpcx.ClientConfigs    `mapstructure:"grpc_client" json:"grpc_client" yaml:"grpc_client"`
	Trace         *otelx.TraceConfig     `mapstructure:"trace" json:"trace" yaml:"trace"`
	Metric        *otelx.MetricConfig    `mapstructure:"metric" json:"metric" yaml:"metric"`
	AMQP          amqpx.Configs          `mapstructure:"amqp" json:"amqp" yaml:"amqp"`
	Kafka         kafkax.Configs         `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Redis         redisx.Configs         `mapstructure:"redis" json:"redis" yaml:"redis"`
	Database      databasex.Configs      `mapstructure:"database" json:"database" yaml:"database"`
	Mongo         mongox.Configs         `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	ElasticSearch elasticsearchx.Configs `mapstructure:"elastic_search" json:"elastic_search" yaml:"elastic_search"`
	Consul        consulx.Configs        `mapstructure:"consul" json:"consul" yaml:"consul"`
	ETCD          etcdx.Configs          `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}

func Env() *EnvConfig {
	return AppConf.Env
}

func Nacos() nacosx.Configs {
	return AppConf.Nacos
}

func GrpcServer() *grpcx.ServerConfig {
	return AppConf.GrpcServer
}

func GrpcClient() grpcx.ClientConfigs {
	return AppConf.GrpcClient
}

func Actuator() *actuatorx.Config {
	return AppConf.Actuator
}

func Kafka() kafkax.Configs {
	return AppConf.Kafka
}

func AMQP() amqpx.Configs {
	return AppConf.AMQP
}

func Trace() *otelx.TraceConfig {
	return AppConf.Trace
}

func Metric() *otelx.MetricConfig {
	return AppConf.Metric
}

func Redis() redisx.Configs {
	return AppConf.Redis
}

func Database() databasex.Configs {
	return AppConf.Database
}

func Mongo() mongox.Configs {
	return AppConf.Mongo
}

func ElasticSearch() elasticsearchx.Configs {
	return AppConf.ElasticSearch
}

func Consul() consulx.Configs {
	return AppConf.Consul
}

func ETCD() etcdx.Configs {
	return AppConf.ETCD
}
