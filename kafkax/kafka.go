package kafkax

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strconv"
	"strings"
)

type Configs map[string]*Config

type Config struct {
	Topic                     string   `mapstructure:"topic" json:"topic" yaml:"topic"`
	Brokers                   []string `mapstructure:"brokers" json:"brokers" yaml:"brokers"`
	GroupID                   string   `mapstructure:"group_id" json:"group_id" yaml:"group_id"`
	SecurityProtocol          string   `mapstructure:"security_protocol" json:"security_protocol" yaml:"security_protocol"`
	SaslMechanism             string   `mapstructure:"sasl_mechanism" json:"sasl_mechanism" yaml:"sasl_mechanism"`
	SaslUsername              string   `mapstructure:"sasl_username" json:"sasl_username" yaml:"sasl_username"`
	SaslPassword              string   `mapstructure:"sasl_password" json:"sasl_password" yaml:"sasl_password"`
	SslCaLocation             string   `mapstructure:"ssl_ca_location" json:"ssl_ca_location" yaml:"ssl_ca_location"`
	EnableSslCertVerification bool     `mapstructure:"enable_ssl_cert_verification" json:"enable_ssl_cert_verification" yaml:"enable_ssl_cert_verification"`
}

func NewConsumers(configs Configs) (map[string]*kafka.Consumer, error) {
	consumers := make(map[string]*kafka.Consumer)
	for key, config := range configs {
		consumer, err := NewConsumer(config)
		if err != nil {
			return nil, err
		}
		consumers[key] = consumer
	}
	return consumers, nil
}

func NewConsumer(config *Config) (*kafka.Consumer, error) {
	configMap := &kafka.ConfigMap{
		"api.version.request":       "true",
		"auto.offset.reset":         "latest",
		"heartbeat.interval.ms":     3000,
		"session.timeout.ms":        30000,
		"max.poll.interval.ms":      120000,
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000,
		"bootstrap.servers":         strings.Join(config.Brokers, ","),
		"group.id":                  config.GroupID,
	}
	setSecurityConfig(config, configMap)
	return kafka.NewConsumer(configMap)
}

func NewProducers(configs Configs) (map[string]*kafka.Producer, error) {
	producers := make(map[string]*kafka.Producer)
	for key, config := range configs {
		producer, err := NewProducer(config)
		if err != nil {
			return nil, err
		}
		producers[key] = producer
	}
	return producers, nil
}

func NewProducer(config *Config) (*kafka.Producer, error) {
	configMap := &kafka.ConfigMap{
		"api.version.request":           "true",
		"message.max.bytes":             1000000,
		"linger.ms":                     500,
		"sticky.partitioning.linger.ms": 1000,
		"retries":                       10,
		"retry.backoff.ms":              1000,
		"acks":                          "1",
		"bootstrap.servers":             strings.Join(config.Brokers, ","),
	}
	setSecurityConfig(config, configMap)
	return kafka.NewProducer(configMap)
}

func setSecurityConfig(config *Config, configMap *kafka.ConfigMap) {
	switch strings.ToLower(config.SecurityProtocol) {
	case "", "plaintext":
		configMap.SetKey("security.protocol", "plaintext")
	case "sasl_ssl":
		configMap.SetKey("security.protocol", "sasl_ssl")
		configMap.SetKey("ssl.ca.location", config.SslCaLocation)
		configMap.SetKey("sasl.username", config.SaslUsername)
		configMap.SetKey("sasl.password", config.SaslPassword)
		configMap.SetKey("sasl.mechanism", config.SaslMechanism)
		configMap.SetKey("enable.ssl.certificate.verification", strconv.FormatBool(config.EnableSslCertVerification))
	case "sasl_plaintext":
		configMap.SetKey("security.protocol", "sasl_plaintext")
		configMap.SetKey("sasl.username", config.SaslUsername)
		configMap.SetKey("sasl.password", config.SaslPassword)
		configMap.SetKey("sasl.mechanism", config.SaslMechanism)
	}
}
