
gen-hms: api/hive_metastore.thrift
	@thrift -r --gen go --gen go:package_prefix=github.com/ozkatz/hive_proxy/pkg/hive/generated/gen-go/ -o pkg/hive/generated api/hive_metastore.thrift
