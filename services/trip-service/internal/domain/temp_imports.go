//go:build ignore
// +build ignore

package main

import (
	_ "go.opentelemetry.io/auto/sdk"
	_ "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	_ "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	_ "go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel/exporters/jaeger"
	_ "go.opentelemetry.io/otel/metric"
	_ "go.opentelemetry.io/otel/sdk"
	_ "go.opentelemetry.io/otel/trace"
	_ "golang.org/x/crypto"
	_ "golang.org/x/net"
	_ "golang.org/x/sync"
	_ "golang.org/x/sys"
	_ "golang.org/x/text"
	_ "google.golang.org/genproto/googleapis/rpc"
	_ "google.golang.org/grpc"
	_ "google.golang.org/protobuf/proto"
)
