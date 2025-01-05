package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"cloud.google.com/go/firestore"

	accountmgrv1 "github.com/t-ash0410/stack-example/go/api/accountmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/account/internal/mgr"
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
		portNumber = os.Getenv("ACCOUNT_MGR_PORT")
		port       = ":" + portNumber

		prjID = os.Getenv("FIRESTORE_PROJECT_ID")
	)

	fsc, err := firestore.NewClient(ctx, prjID)
	if err != nil {
		return fmt.Errorf("failed to create firestore client: %w", err)
	}

	s := grpcx.NewServer()
	server, err := mgr.NewAccountMgrServer(fsc)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	accountmgrv1.RegisterAccountMgrServiceServer(s.Srv, server)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	return s.Run(ctx, lis)
}
