package plugin

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/logical"
)

type accessConfig struct {
	Address       string `json:"address"`
	CACert        string `json:"ca_certificate"`
	ClientCert    string `json:"client_certificate"`
	ClientKey     string `json:"client_key"`
	TLSEnabled    bool   `json:"tls_enabled"`
	SkipTLSVerify bool   `json:"skip_tls_verify"`
}

func (a *accessConfig) cert() (tls.Certificate, error) {
	return tls.X509KeyPair([]byte(a.ClientCert), []byte(a.ClientKey))
}

func client(ctx context.Context, s logical.Storage) (*kafka.Client, error, error) {
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
	}

	log.Printf("[DEBUG] Client certicate being loaded")
	cert, err := conf.cert()
	if err == nil {
		kafkaConfig.ClientCert = &cert
	} else {
		return nil, err, nil
	}

	log.Printf("[DEBUG] CA certicate being loaded")
	block, _ := pem.Decode([]byte(conf.CACert))
	var cacert *x509.Certificate
	cacert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err, nil
	}
	if cacert != nil {
		kafkaConfig.CACert = cacert
	}

	log.Printf("[DEBUG] Creating Client")
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
		return nil, nil, errwrap.Wrapf("error reading consul access configuration: {{err}}", err)
	}

	return conf, nil, nil
}
