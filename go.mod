module github.com/rubiskko/greeting-grpc

go 1.20

replace (
	github.com/rubiskko/greeting-grpc/api => ./api
	github.com/rubiskko/greeting-grpc/cli => ./cli
	github.com/rubiskko/greeting-grpc/server => ./server
)

require (
	github.com/rubiskko/greeting-grpc/cli v0.0.0-00010101000000-000000000000
	github.com/rubiskko/greeting-grpc/server v0.0.0-00010101000000-000000000000
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rubiskko/greeting-grpc/api v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.29.0 // indirect
)
