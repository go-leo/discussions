package configx

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	Env,
	Nacos,
	Gin,
	GrpcServer,
	GrpcClient,
	Actuator,
	Kafka,
	AMQP,
	Trace,
	Metric,
	Redis,
	Database,
	Mongo,
	ElasticSearch,
	Consul,
	ETCD,
)
