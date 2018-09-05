package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"runtime"
)

const (
	version     = "0.2.0"
	defaultAddr = ":7887"
)

var GitHash = "unknown"

type options struct {
	Version     bool   `short:"v" long:"version" description:"Print version info."`
	Addr        string `short:"a" long:"addr" description:"IP address to listen on. Default: ':7887'." env:"FAKESM_ADDR"`
	SecretsJson string `short:"f" long:"secrets-json" description:"Path to JSON file containing the secrets in a hash. The JSON hash key names are the secret names. If the secret is itself JSON, it needs to be escaped and stuffed in there as a normal string." env:"FAKESM_SECRETS_JSON"`
}

func parseOptions() (*options, error) {
	var opts = &options{}
	parser := flags.NewParser(opts, flags.Default)
	parser.ShortDescription = fmt.Sprintf("A stand-in for AWS Secrets Manager for local testing that doesn't require access to the real Secrets Manager - version %s", version)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			return nil, err
		}
	}
	if opts.Version {
		fmt.Printf("fake-secretsmanager %s (git hash: %s) built with %s.\n", version, GitHash, runtime.Version())
		os.Exit(0)
	}
	if opts.Addr == "" {
		opts.Addr = defaultAddr
	}
	if opts.SecretsJson == "" {
		return nil, fmt.Errorf("No secrets file provided!")
	}

	return opts, nil
}
