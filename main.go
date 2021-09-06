package main

import (
	"log"
	"os"

	k "github.com/Mongey/vault-plugin-secrets-kafka/plugin"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: k.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})

	if err != nil {
		log.Println(err)

		os.Exit(1)
	}
}
