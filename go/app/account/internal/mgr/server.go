package mgr

import (
	"fmt"

	"cloud.google.com/go/firestore"

	accountmgrv1 "github.com/t-ash0410/stack-example/go/api/accountmgr/v1"
)

type AccountMgrServer struct {
	fsc *firestore.Client

	accountmgrv1.UnimplementedAccountMgrServiceServer
}

func NewAccountMgrServer(fsc *firestore.Client) (*AccountMgrServer, error) {
	if fsc == nil {
		return nil, fmt.Errorf("invalid arguments")
	}
	return &AccountMgrServer{
		fsc: fsc,
	}, nil
}
