module github.com/Casper-Mars/cloud-restaurant/user

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/Casper-Mars/cloud-restaurant v0.0.0-20210819063323-2633caede073
	github.com/go-kratos/kratos/v2 v2.0.5
	github.com/google/wire v0.5.0
	github.com/lib/pq v1.10.2
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	google.golang.org/protobuf v1.27.1
)

//
//replace (
//	github.com/Casper-Mars/cloud-restaurant  => ../../
//)
