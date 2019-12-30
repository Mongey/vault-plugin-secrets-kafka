package plugin

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathConfigAccess() *framework.Path {
	return &framework.Path{
		Pattern: "config/access",
		Fields: map[string]*framework.FieldSchema{
			"address": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Kafka server address",
			},

			"ca_certificate": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "The certificate ca",
			},
			"client_certificate": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "The certificate of the client. This should probably be a superruser.",
			},
			"client_key": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "The key for the client certificate",
			},
			"tls_enabled": &framework.FieldSchema{
				Type:        framework.TypeBool,
				Description: "Enable TLS for communication with the broker",
				Default:     true,
			},
			"skip_tls_verify": &framework.FieldSchema{
				Type:        framework.TypeBool,
				Description: "Skip verification of the connection with the broker",
				Default:     false,
				//Optional:    true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation:   pathConfigAccessRead,
			logical.UpdateOperation: pathConfigAccessWrite,
		},
	}
}

func pathConfigAccessRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	conf, userErr, intErr := readConfigAccess(ctx, req.Storage)
	if intErr != nil {
		return nil, intErr
	}
	if userErr != nil {
		return logical.ErrorResponse(userErr.Error()), nil
	}
	if conf == nil {
		return nil, fmt.Errorf("no user error reported but consul access configuration not found")
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"address":            conf.Address,
			"ca_certificate":     conf.CACert,
			"client_certificate": conf.ClientCert,
			"client_key":         conf.ClientKey,
			"tls_enabled":        conf.TLSEnabled,
			"skip_tls_verify":    conf.SkipTLSVerify,
		},
	}, nil
}

func pathConfigAccessWrite(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	entry, err := logical.StorageEntryJSON("config/access", accessConfig{
		Address:       data.Get("address").(string),
		CACert:        data.Get("ca_certificate").(string),
		ClientCert:    data.Get("client_certificate").(string),
		ClientKey:     data.Get("client_key").(string),
		TLSEnabled:    data.Get("tls_enabled").(bool),
		SkipTLSVerify: data.Get("skip_tls_verify").(bool),
	})

	if err != nil {
		return nil, err
	}

	if err := req.Storage.Put(ctx, entry); err != nil {
		return nil, err
	}

	log.Printf("[INFO] ALL WENT OK")
	return nil, nil
}
