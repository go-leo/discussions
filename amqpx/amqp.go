package amqpx

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/rabbitmq/amqp091-go"
	"os"
)

type Configs map[string]*Config

type Config struct {
	URL           string          `mapstructure:"url" json:"url" yaml:"url"`
	TLS           *TLSConfig      `mapstructure:"tls" json:"tls" yaml:"tls"`
	ExternalAuth  bool            `mapstructure:"external_auth" json:"external_auth" yaml:"external_auth"`
	Amqp091Config *amqp091.Config `mapstructure:"amqp_091_config" json:"amqp_091_config" yaml:"amqp_091_config"`
}

type TLSConfig struct {
	PEMFile  string `mapstructure:"pem_file" json:"pem_file" yaml:"pem_file"`
	CertFile string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file"`
	KeyFile  string `mapstructure:"key_file" json:"key_file" yaml:"key_file"`
}

func NewConnections(configs Configs) (map[string]*amqp091.Connection, error) {
	connections := make(map[string]*amqp091.Connection)
	for key, config := range configs {
		connection, err := NewConnection(config)
		if err != nil {
			return nil, err
		}
		connections[key] = connection
	}
	return connections, nil
}

func NewConnection(config *Config) (*amqp091.Connection, error) {
	if len(config.URL) <= 0 {
		return nil, errors.New("url is empty")
	}
	if config.Amqp091Config != nil {
		return amqp091.DialConfig(config.URL, *config.Amqp091Config)
	}
	if config.TLS == nil {
		return amqp091.Dial(config.URL)
	}
	tlsCfg := new(tls.Config)
	tlsCfg.RootCAs = x509.NewCertPool()
	if ca, err := os.ReadFile(config.TLS.PEMFile); err == nil {
		tlsCfg.RootCAs.AppendCertsFromPEM(ca)
	}
	if cert, err := tls.LoadX509KeyPair(config.TLS.CertFile, config.TLS.KeyFile); err == nil {
		tlsCfg.Certificates = append(tlsCfg.Certificates, cert)
	}
	if config.ExternalAuth {
		return amqp091.DialTLS_ExternalAuth(config.URL, tlsCfg)
	}
	return amqp091.DialTLS(config.URL, tlsCfg)
}

//
//type ExchangeConfig struct {
//	Name       string
//	Kind       string
//	Durable    bool
//	AutoDelete bool
//	Internal   bool
//	NoWait     bool
//	Args       map[string]interface{}
//}
//
//func ExchangeDeclare(channel *amqp091.Channel, config *ExchangeConfig) error {
//	return channel.ExchangeDeclare(
//		config.Name,
//		config.Kind,
//		config.Durable,
//		config.AutoDelete,
//		config.Internal,
//		config.NoWait,
//		config.Args,
//	)
//}
//
//type QueueConfig struct {
//	Name       string
//	Durable    bool
//	AutoDelete bool
//	Exclusive  bool
//	NoWait     bool
//	Args       map[string]interface{}
//}
//
//func QueueDeclare(channel *amqp091.Channel, config *QueueConfig) (amqp091.Queue, error) {
//	return channel.QueueDeclare(config.Name, config.Durable, config.AutoDelete, config.Exclusive, config.NoWait, config.Args)
//}
//
//type QueueBindConfig struct {
//	Key      string
//	Exchange string
//	NoWait   bool
//	Args     map[string]interface{}
//}
//
//func QueueBind(channel *amqp091.Channel, queue amqp091.Queue, config *QueueBindConfig) error {
//	return channel.QueueBind(queue.Name, config.Key, config.Exchange, config.NoWait, config.Args)
//}
//
//type QosConfig struct {
//	PrefetchCount int
//	PrefetchSize  int
//	Global        bool
//}
//
//func EqualizeMessageDistribution(channel *amqp091.Channel, config *QosConfig) error {
//	return channel.Qos(config.PrefetchCount, config.PrefetchSize, config.Global)
//}
//
//func Ack(channel *amqp091.Channel, tag uint64, multiple bool) error {
//	return channel.Ack(tag, multiple)
//}
//
//func Nack(channel *amqp091.Channel, tag uint64, multiple bool, requeue bool) error {
//	return channel.Nack(tag, multiple, requeue)
//}
