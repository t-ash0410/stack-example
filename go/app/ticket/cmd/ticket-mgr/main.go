package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"cloud.google.com/go/firestore"

	ticketmgrv1 "github.com/t-ash0410/stack-example/go/api/ticketmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/mgr"
	"github.com/t-ash0410/stack-example/go/lib/grpcx"
)

const (
	exitFail = 1
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(exitFail)
	}
}

func run(ctx context.Context) error {
	var (
		portNumber = os.Getenv("TICKET_MGR_PORT")
		port       = ":" + portNumber

		prjID = os.Getenv("FIRESTORE_PROJECT_ID")
	)

	fsc, err := firestore.NewClient(ctx, prjID)
	if err != nil {
		return fmt.Errorf("failed to create firestore client: %w", err)
	}

	s := grpcx.NewServer()
	server, err := mgr.NewTicketMgrServer(fsc)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	ticketmgrv1.RegisterTicketMgrServiceServer(s.Srv, server)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	return s.Run(ctx, lis)
}
