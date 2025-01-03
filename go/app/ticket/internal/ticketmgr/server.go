package ticketmgr

import (
	"fmt"

	"cloud.google.com/go/firestore"

	ticketmgrv1 "github.com/t-ash0410/stack-example/go/api/ticketmgr/v1"
)

type TicketMgrServer struct {
	fsc *firestore.Client

	ticketmgrv1.UnimplementedTicketMgrServiceServer
}

func NewTicketMgrServer(fsc *firestore.Client) (*TicketMgrServer, error) {
	if fsc == nil {
		return nil, fmt.Errorf("invalid arguments")
	}
	return &TicketMgrServer{
		fsc: fsc,
	}, nil
}
