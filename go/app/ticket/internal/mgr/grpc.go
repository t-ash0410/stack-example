package mgr

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ticketmgrv1 "github.com/t-ash0410/stack-example/go/api/ticketmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/modelfs"
	"github.com/t-ash0410/stack-example/go/lib/firestorex"
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
		return nil, status.Errorf(codes.Internal, "failed to create ticket: %v", err)
	}

	return &ticketmgrv1.CreateTicketResponse{
		TicketId: t.TicketID,
	}, nil
}

func newTicketFromCreateReq(req *ticketmgrv1.CreateTicketRequest) (*modelfs.Ticket, error) {
	var (
		t = &modelfs.Ticket{
			TicketID:    uuid.NewString(),
			Title:       req.Title,
			CreatedBy:   req.RequestedBy,
			UpdatedBy:   req.RequestedBy,
			Description: req.Description,
			Deadline:    req.Deadline.AsTime(),
		}
	)
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TicketMgrServer) UpdateTicket(ctx context.Context,
	req *ticketmgrv1.UpdateTicketRequest,
) (*ticketmgrv1.UpdateTicketResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	err := s.fsc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		ref := s.fsc.Doc(fmt.Sprintf("%s/%s", modelfs.CollectionNameTickets, req.TicketId))
		t, err := firestorex.ReadOneWithTxn[modelfs.Ticket](tx, ref)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to read, ticket id = %q: %v", req.TicketId, err)
		}

		if err := updateTicketByUpdateReq(t.Data, req); err != nil {
			return status.Errorf(codes.InvalidArgument, "failed to validate, ticket id = %q: %v", req.TicketId, err)
		}

		return tx.Set(ref, t.Data)
	})
	if err != nil {
		return nil, err
	}

	return &ticketmgrv1.UpdateTicketResponse{}, nil
}

func updateTicketByUpdateReq(t *modelfs.Ticket, req *ticketmgrv1.UpdateTicketRequest) error {
	t.UpdatedBy = req.RequestedBy

	if req.Title != nil {
		t.Title = req.GetTitle()
	}
	if req.Description != nil {
		t.Description = req.GetDescription()
	}
	if req.Deadline != nil {
		t.Deadline = req.GetDeadline().AsTime()
	}

	return t.Validate()
}

func (s *TicketMgrServer) DeleteTicket(ctx context.Context,
	req *ticketmgrv1.DeleteTicketRequest,
) (*ticketmgrv1.DeleteTicketResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	ref := s.fsc.Doc(fmt.Sprintf("%s/%s", modelfs.CollectionNameTickets, req.TicketId))
	err := s.fsc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return tx.Delete(ref)
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete ticket, ticket id = %q: %v", req.TicketId, err)
	}

	return &ticketmgrv1.DeleteTicketResponse{}, nil
}
