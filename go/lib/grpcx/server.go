package grpcx

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Srv *grpc.Server
}

func NewServer() *Server {
	var (
		s  = grpc.NewServer()
		hs = health.NewServer()
	)
	reflection.Register(s)
	grpc_health_v1.RegisterHealthServer(s, hs)

	return &Server{
		Srv: s,
	}
}

func (s *Server) Run(ctx context.Context, l net.Listener) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		if err := s.Srv.Serve(l); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("failed to serve: %w", err)
		}
	case <-ctx.Done():
		s.Srv.GracefulStop()
	}

	return nil
}
