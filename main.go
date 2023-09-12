package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/ozkatz/hive_proxy/pkg/hive"
)

const (
	DefaultListenAddr       = "127.0.0.1:9083"
	DefaultHiveMetastoreURI = "thrift://127.0.0.1:9093"
)

func dieIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func die(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDRESS")
	if listenAddr == "" {
		listenAddr = DefaultListenAddr
	}
	metastoreUri := os.Getenv("HIVE_METASTORE_URI")
	if metastoreUri == "" {
		metastoreUri = DefaultHiveMetastoreURI
	}
	parsedUri, err := url.Parse(metastoreUri)
	dieIfErr(err)
	if parsedUri.Scheme != "thrift" {
		die("invalid uri: %s, expected thrift URI\n", metastoreUri)
	}

	fmt.Printf("Connecting to: %s\nListening on: thrift://%s\n\n", parsedUri.Host, listenAddr)
	ctx := context.Background()
	if err = hive.RunProxyServer(ctx, parsedUri.Host, listenAddr); err != nil {
		die("Error running proxy: %+v\n", err)
	}
}
