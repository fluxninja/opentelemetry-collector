module go.opentelemetry.io/collector

go 1.20

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.2
	github.com/prometheus/client_golang v1.17.0
	github.com/prometheus/client_model v0.5.0
	github.com/prometheus/common v0.44.0
	github.com/shirou/gopsutil/v3 v3.23.9
	github.com/stretchr/testify v1.8.4
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector/component v0.88.0
	go.opentelemetry.io/collector/config/configtelemetry v0.88.0
	go.opentelemetry.io/collector/connector v0.88.0
	go.opentelemetry.io/collector/consumer v0.88.0
	go.opentelemetry.io/collector/exporter v0.88.0
	go.opentelemetry.io/collector/featuregate v1.0.0-rcv0017
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0017
	go.opentelemetry.io/collector/processor v0.88.0
	go.opentelemetry.io/collector/receiver v0.88.0
	go.opentelemetry.io/collector/service v0.88.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/prometheus v0.42.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.opentelemetry.io/otel/sdk/metric v1.19.0
	go.uber.org/multierr v1.11.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/prometheus/statsd_exporter v0.22.8 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.opentelemetry.io/collector/confmap v0.88.0 // indirect
	go.opentelemetry.io/collector/extension v0.88.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	go.uber.org/goleak v1.2.1 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.opentelemetry.io/collector/component => ./component

replace go.opentelemetry.io/collector/confmap => ./confmap

replace go.opentelemetry.io/collector/config/confignet => ./config/confignet

replace go.opentelemetry.io/collector/config/configtelemetry => ./config/configtelemetry

replace go.opentelemetry.io/collector/connector => ./connector

replace go.opentelemetry.io/collector/consumer => ./consumer

replace go.opentelemetry.io/collector/exporter => ./exporter

replace go.opentelemetry.io/collector/extension => ./extension

replace go.opentelemetry.io/collector/featuregate => ./featuregate

replace go.opentelemetry.io/collector/semconv => ./semconv

replace go.opentelemetry.io/collector/pdata => ./pdata

replace go.opentelemetry.io/collector/processor => ./processor

replace go.opentelemetry.io/collector/receiver => ./receiver

replace go.opentelemetry.io/collector/service => ./service

replace go.opentelemetry.io/collector/extension/zpagesextension => ./extension/zpagesextension

replace go.opentelemetry.io/collector/processor/batchprocessor => ./processor/batchprocessor

retract (
	v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module, use v0.76.1
	v0.69.0 // Release failed, use v0.69.1
	v0.57.1 // Release failed, use v0.57.2
	v0.57.0 // Release failed, use v0.57.2
	v0.32.0 // Contains incomplete metrics transition to proto 0.9.0, random components are not working.
)
