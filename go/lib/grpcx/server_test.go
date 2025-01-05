package grpcx_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"

	"github.com/t-ash0410/stack-example/go/lib/grpcx"
)

func TestNewServer(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		s := grpcx.NewServer()
		assert.NotNil(t, s)
	})
}

func TestServer_Run(t *testing.T) {
	t.Parallel()

	t.Run("Success: Termination by context cancellation", func(t *testing.T) {
		t.Parallel()

		var (
			ctx, cancel = context.WithCancel(context.Background())
			s           = grpcx.NewServer()
		)
		cancel()

		lis, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatalf("Failed to listen: %v", err)
		}

		err = s.Run(ctx, lis)
		assert.Nil(t, err)
	})

	t.Run("Fail: Termination by closed connection", func(t *testing.T) {
		t.Parallel()

		var (
			ctx, cancel = context.WithCancel(context.Background())
			s           = grpcx.NewServer()
		)
		defer cancel()

		lis, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatalf("Failed to listen: %v", err)
		}
		lis.Close()

		err = s.Run(ctx, lis)
		assert.ErrorIs(t, err, net.ErrClosed)
	})

	t.Run("The gRPC service to be used in common must be registered", func(t *testing.T) {
		t.Parallel()

		var (
			ctx, cancel = context.WithCancel(context.Background())
			s           = grpcx.NewServer()
		)
		t.Cleanup(cancel)

		lis, err := net.Listen("tcp", ":19080")
		if err != nil {
			t.Fatalf("Failed to listen: %v", err)
		}

		go s.Run(ctx, lis)

		time.Sleep(10 * time.Millisecond)

		c, err := grpc.NewClient("localhost:19080", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Fatalf("Failed to create gRPC client: %v", err)
		}

		src := grpc_reflection_v1.NewServerReflectionClient(c)
		srRes, err := src.ServerReflectionInfo(ctx)
		assert.NotNil(t, srRes)
		assert.NoError(t, err)

		hc := grpc_health_v1.NewHealthClient(c)
		hcRes, err := hc.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		assert.NotNil(t, hcRes)
		assert.NoError(t, err)

		cancel()
	})
}
