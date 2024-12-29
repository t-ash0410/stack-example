package ticketmgr

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ticketmgrv1 "github.com/t-ash0410/stack-example/go/api/ticketmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/firestorex"
)

func (s *TicketMgrServer) CreateTicket(ctx context.Context,
	req *ticketmgrv1.CreateTicketRequest,
) (*ticketmgrv1.CreateTicketResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	t, err := newTicketFromCreateReq(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to validate ticket: %v", err)
	}

	err = s.fsc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return tx.Create(s.fsc.Doc(t.Path()), t)
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create a document: %v", err)
	}

	return &ticketmgrv1.CreateTicketResponse{
		TicketId: t.TicketID,
	}, nil
}

func newTicketFromCreateReq(req *ticketmgrv1.CreateTicketRequest) (*firestorex.Ticket, error) {
	var (
		now = time.Now()
		t   = &firestorex.Ticket{
			TicketID:    uuid.NewString(),
			Title:       req.Title,
			CreatedBy:   req.RequestedBy,
			CreatedAt:   now,
			UpdatedAt:   now,
			Description: req.Description,
			Deadline:    req.Deadline.AsTime(),
		}
	)
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}
