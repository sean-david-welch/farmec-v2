module github.com/sean-david-welch/farmec-v2/server

go 1.23.3

require (
	firebase.google.com/go v3.13.0+incompatible
	github.com/aws/aws-sdk-go-v2 v1.36.1
	github.com/aws/aws-sdk-go-v2/config v1.29.6
	github.com/aws/aws-sdk-go-v2/credentials v1.17.59
	github.com/aws/aws-sdk-go-v2/service/s3 v1.76.0
	github.com/aws/smithy-go v1.22.2
	github.com/gin-contrib/cors v1.7.3
	github.com/gin-gonic/gin v1.10.0
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/sendgrid/sendgrid-go v3.16.0+incompatible
	github.com/signintech/gopdf v0.29.2
	github.com/stripe/stripe-go/v76 v76.25.0
	google.golang.org/api v0.220.0
)

require (
	cel.dev/expr v0.19.0 // indirect
	cloud.google.com/go v0.117.0 // indirect
	cloud.google.com/go/auth v0.14.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.7 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/firestore v1.18.0 // indirect
	cloud.google.com/go/iam v1.2.2 // indirect
	cloud.google.com/go/longrunning v0.6.2 // indirect
	cloud.google.com/go/monitoring v1.21.2 // indirect
	cloud.google.com/go/storage v1.50.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.25.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.48.1 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.48.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.5.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.14 // indirect
	github.com/bytedance/sonic v1.12.6 // indirect
	github.com/bytedance/sonic/loader v0.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/envoyproxy/go-control-plane/envoy v1.32.3 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.23.0 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/phpdave11/gofpdi v1.0.14-0.20211212211723-1f10f9844311 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/sendgrid/rest v2.6.9+incompatible // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.32.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.58.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.58.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	golang.org/x/arch v0.12.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.26.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250127172529-29210b9bc287 // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
