module go.opentelemetry.io/collector/exporter

go 1.20

require (
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/fluxninja/aperture v1.6.0
	github.com/stretchr/testify v1.8.4
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector v0.91.0
	go.opentelemetry.io/collector/component v0.91.0
	go.opentelemetry.io/collector/config/configtelemetry v0.91.0
	go.opentelemetry.io/collector/consumer v0.91.0
	go.opentelemetry.io/collector/extension v0.91.0
	go.opentelemetry.io/collector/featuregate v1.0.0
	go.opentelemetry.io/collector/pdata v1.0.0
	go.opentelemetry.io/collector/receiver v0.91.0
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
	golang.org/x/sys v0.15.0
	google.golang.org/grpc v1.59.0
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/prometheus/statsd_exporter v0.22.8 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	go.opentelemetry.io/collector/confmap v0.91.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.44.1-0.20231201153405-6027c1ae76f2 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.21.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.opentelemetry.io/collector => ../

replace go.opentelemetry.io/collector/component => ../component

replace go.opentelemetry.io/collector/confmap => ../confmap

replace go.opentelemetry.io/collector/consumer => ../consumer

replace go.opentelemetry.io/collector/extension => ../extension

replace go.opentelemetry.io/collector/featuregate => ../featuregate

replace go.opentelemetry.io/collector/pdata => ../pdata

replace go.opentelemetry.io/collector/processor => ../processor

replace go.opentelemetry.io/collector/receiver => ../receiver

retract v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module

replace go.opentelemetry.io/collector/connector => ../connector

replace go.opentelemetry.io/collector/config/configtelemetry => ../config/configtelemetry
