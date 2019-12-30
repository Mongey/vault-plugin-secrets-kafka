package plugin

import (
	"context"
	"fmt"
	"log"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/Shopify/sarama"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/logical"
)

type accessConfig struct {
	Address       string `json:"address"`
	CACert        string `json:"ca_certificate"`
	ClientCert    string `json:"client_certificate"`
	ClientKey     string `json:"client_key"`
	TLSEnabled    bool   `json:"tls_enabled"`
	SkipTLSVerify bool   `json:"skip_tls_verify"`
}

func client(ctx context.Context, s logical.Storage) (*kafka.Client, error, error) {
	logger := hclog.New(&hclog.LoggerOptions{})
	sarama.Logger = log.New(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true, ForceLevel: hclog.Trace}), "[DEBUG] [Sarama]", log.LstdFlags)
	conf, userErr, intErr := readConfigAccess(ctx, s)
	if intErr != nil {
		return nil, nil, intErr
	}
	if userErr != nil {
		return nil, userErr, nil
	}
	if conf == nil {
		return nil, nil, fmt.Errorf("no error received but no configuration found")
	}

	kafkaConfig := &kafka.Config{
		BootstrapServers: &[]string{conf.Address},
		TLSEnabled:       conf.TLSEnabled,
		SkipTLSVerify:    conf.SkipTLSVerify,
		ClientCert:       conf.ClientCert,
		ClientCertKey:    conf.ClientKey,
		CACert:           conf.CACert,
		Timeout:          300,
	}

	log.Printf("[DEBUG] Creating Client %s", conf.Address)
	c, err := kafka.NewClient(kafkaConfig)
	return c, nil, err
}

func readConfigAccess(ctx context.Context, storage logical.Storage) (*accessConfig, error, error) {
	entry, err := storage.Get(ctx, "config/access")
	if err != nil {
		return nil, nil, err
	}

	if entry == nil {
		return nil, fmt.Errorf("access credentials for the backend itself haven't been configured; please configure them at the '/config/access' endpoint"), nil
	}

	conf := &accessConfig{}
	if err := entry.DecodeJSON(conf); err != nil {
		return nil, nil, fmt.Errorf("error reading consul access configuration: %s", err)
	}

	return conf, nil, nil
}
