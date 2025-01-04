package querier

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	ticketquerierv1 "github.com/t-ash0410/stack-example/go/api/ticketquerier/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/modelfs"
	"github.com/t-ash0410/stack-example/go/lib/firestorex"
)

func (s *TicketQuerierServer) QueryTickets(ctx context.Context,
	req *ticketquerierv1.QueryTicketsRequest,
) (*ticketquerierv1.QueryTicketsResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var (
		res  = &ticketquerierv1.QueryTicketsResponse{}
		iter = s.fsc.Collection(modelfs.CollectionNameTickets).Where("CreatedBy", "==", req.RequestedBy).Documents(ctx)
	)
	for d, err := range firestorex.ReadEach[modelfs.Ticket](iter) {
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to read ticket: %v", err)
		}
		res.Tickets = append(res.Tickets, ticketModelToPB(d))
	}

	return res, nil
}

func ticketModelToPB(d *firestorex.ResultWithMeta[modelfs.Ticket]) *ticketquerierv1.Ticket {
	return &ticketquerierv1.Ticket{
		TicketId:    d.Data.TicketID,
		CreatedAt:   timestamppb.New(d.CreateTime),
		UpdatedAt:   timestamppb.New(d.UpdateTime),
		CreatedBy:   d.Data.CreatedBy,
		Title:       d.Data.Title,
		Description: d.Data.Description,
		Deadline:    timestamppb.New(d.Data.Deadline),
	}
}
