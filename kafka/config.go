package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Config struct {
	BootstrapServers        *[]string
	Timeout                 int
	CACert                  *x509.Certificate
	ClientCert              *tls.Certificate
	ClientCertKey           string
	ClientCertKeyPassphrase string
	TLSEnabled              bool
	SkipTLSVerify           bool
	SASLUsername            string
	SASLPassword            string
	SASLMechanism           string
}

func (c *Config) newKafkaConfig() (*sarama.Config, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = sarama.V2_4_0_0
	kafkaConfig.ClientID = "vault-plugin-secrets-kafka"
	kafkaConfig.Admin.Timeout = time.Duration(c.Timeout) * time.Second
	kafkaConfig.Metadata.Full = true // the default, but just being clear

	if c.saslEnabled() {
		switch c.SASLMechanism {
		case "scram-sha512":
			kafkaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
			kafkaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA512)
		case "scram-sha256":
			kafkaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
			kafkaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA256)
		case "plain":
		default:
			log.Fatalf("[ERROR] Invalid sasl mechanism \"%s\": can only be \"scram-sha256\", \"scram-sha512\" or \"plain\"", c.SASLMechanism)
		}
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.Password = c.SASLPassword
		kafkaConfig.Net.SASL.User = c.SASLUsername
		kafkaConfig.Net.SASL.Handshake = true
	} else {
		log.Printf("[WARN] SASL disabled username: '%s', password '%s'", c.SASLUsername, "****")
	}

	if c.TLSEnabled {
		tlsConfig, err := newTLSConfig(
			c.ClientCert,
			c.ClientCertKey,
			c.CACert,
			c.ClientCertKeyPassphrase,
		)

		if err != nil {
			return kafkaConfig, err
		}

		kafkaConfig.Net.TLS.Config = tlsConfig
		kafkaConfig.Net.TLS.Enable = true
		kafkaConfig.Net.TLS.Config.InsecureSkipVerify = c.SkipTLSVerify
	}

	return kafkaConfig, nil
}

func (c *Config) saslEnabled() bool {
	return c.SASLUsername != "" || c.SASLPassword != ""
}

func newTLSConfig(clientCert *tls.Certificate, clientKey string, caCert *x509.Certificate, clientKeyPassphrase string) (*tls.Config, error) {
	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{*clientCert},
	}

	return &tlsConfig, nil
}
