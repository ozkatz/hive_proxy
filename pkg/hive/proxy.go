package hive

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	tutil "github.com/ozkatz/thrift-utils"

	"github.com/ozkatz/hive_proxy/pkg/hive/generated/gen-go/hive_metastore"
)

func connectUpstream(addr string) (*hive_metastore.ThriftHiveMetastoreClient, error) {
	connectFn := func() (thrift.TClient, error) {
		cfg := &thrift.TConfiguration{}
		transport := thrift.NewTSocketConf(addr, cfg)
		if err := transport.Open(); err != nil {
			return nil, err
		}
		protocolFactory := thrift.NewTBinaryProtocolFactoryConf(cfg)
		return thrift.NewTStandardClient(
			protocolFactory.GetProtocol(transport),
			protocolFactory.GetProtocol(transport)), nil
	}
	client, err := tutil.NewRertryingClient(
		connectFn, tutil.RetryOnNetError, tutil.DefaultExponentialBackoff)
	if err != nil {
		return nil, err
	}
	return hive_metastore.NewThriftHiveMetastoreClient(client), nil
}

func RunProxyServer(ctx context.Context, upstream, addr string) error {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	client, err := connectUpstream(upstream)
	if err != nil {
		return err
	}

	processor := hive_metastore.NewThriftHiveMetastoreProcessor(client)

	tproc := tutil.Log(processor, func(c *tutil.Call) {
		if c.Err != nil {
			fmt.Printf("name: %s\ninput: %s\noutput: %s\nerror: %s\n\n\n",
				c.Name, c.Input, c.Output, c.Err)
		} else {
			fmt.Printf("name: %s\ninput: %s\noutput: %s\n\n\n",
				c.Name, c.Input, c.Output)
		}

	})
	server := thrift.NewTSimpleServer4(tproc, transport, transportFactory, protocolFactory)
	return server.Serve()
}
