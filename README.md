# Hive Proxy

This is a Thrift proxy server, wrapping an existing Hive Metastore Thrift server, simply to log which calls are exectued, with their inputs and outputs.

## Running a proxy

```shell
git clone https://github.com/ozkatz/hive_proxy
$ cd hive_proxy
$ go build .
$ export HIVE_METASTORE_URI="thrift://127.0.0.1:9093"  # This is the default value
$ export LISTEN_ADDRESS="127.0.0.1:9083"  # Also a default value
$ ./hive_proxy
```

HMS calls made to `127.0.0.1:9083` will be proxied to `127.0.0.1:9093` and printed to stdout.


## License

[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0). 

See [LICENSE](./LICENSE) and [NOTICE](./NOTICE).
