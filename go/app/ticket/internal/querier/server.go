package querier

import (
	"fmt"

	"cloud.google.com/go/firestore"

	ticketquerierv1 "github.com/t-ash0410/stack-example/go/api/ticketquerier/v1"
)

type TicketQuerierServer struct {
	fsc *firestore.Client

	ticketquerierv1.UnimplementedTicketQuerierServiceServer
}

func NewTicketQuerierServer(fsc *firestore.Client) (*TicketQuerierServer, error) {
	if fsc == nil {
		return nil, fmt.Errorf("invalid arguments")
	}
	return &TicketQuerierServer{
		fsc: fsc,
	}, nil
}
