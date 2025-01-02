package ticketmgr_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	ticketmgrv1 "github.com/t-ash0410/stack-example/go/api/ticketmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/firestorex"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/ticketmgr"
	"github.com/t-ash0410/stack-example/go/lib/ctxtest"
	"github.com/t-ash0410/stack-example/go/lib/firestoretest"
)

func TestServer_CreateTicket(t *testing.T) {
	var (
		t2024_12_29_UTC = time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC)

		baseReq = &ticketmgrv1.CreateTicketRequest{
			Title:       "Some Ticket",
			RequestedBy: "8ea79f88-5b4b-4df6-b438-81a2ccf6b09f",
			Description: "Some ticket description.",
			Deadline:    timestamppb.New(t2024_12_29_UTC.AddDate(0, 0, 10)),
		}
	)

	t.Run("Success", func(t *testing.T) {
		cases := map[string]struct {
			req *ticketmgrv1.CreateTicketRequest

			want       *ticketmgrv1.CreateTicketResponse
			wantTicket *firestorex.Ticket
		}{
			"Create a some ticket": {
				req:  baseReq,
				want: &ticketmgrv1.CreateTicketResponse{
					// TicketId
				},
				wantTicket: &firestorex.Ticket{
					// TicketID
					Title:       baseReq.Title,
					CreatedBy:   baseReq.RequestedBy,
					UpdatedBy:   baseReq.RequestedBy,
					Description: baseReq.Description,
					Deadline:    baseReq.Deadline.AsTime(),
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				fsc, err := firestoretest.InitFirestoreClient(context.Background(), "tickets")
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				s, err := ticketmgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				res, err := s.CreateTicket(context.Background(), tc.req)
				if !assert.Nil(t, err) {
					return
				}
				if diff := cmp.Diff(tc.want, res, protocmp.Transform(), protocmp.IgnoreFields(tc.want, "ticket_id")); diff != "" {
					t.Errorf("Response didn't match (-want / +got)\n%s", diff)
				}

				d := readTicket(t, fsc, res.TicketId)
				if diff := cmp.Diff(tc.wantTicket, d, cmpopts.IgnoreFields(*tc.wantTicket, "TicketID")); diff != "" {
					t.Errorf("Stored data didn't match (-want / +got)\n%s", diff)
				}
			})
		}
	})

	t.Run("Fail", func(t *testing.T) {
		cases := map[string]struct {
			ctx context.Context // optional
			req *ticketmgrv1.CreateTicketRequest

			wantErr assert.ErrorAssertionFunc
		}{
			"Context cancelled": {
				ctx: ctxtest.CanceledContext(),
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.EqualError(t, err, context.Canceled.Error())
				},
			},
			"InvalidArgument: Validation error": {
				req: &ticketmgrv1.CreateTicketRequest{
					RequestedBy: "",  // important
					Title:       "",  // important
					Description: "",  // important
					Deadline:    nil, // important
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Errorf(codes.InvalidArgument, "failed to validate ticket").Error())
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				fsc, err := firestoretest.InitFirestoreClient(context.Background(), "tickets")
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				s, err := ticketmgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				ctx := context.Background()
				if tc.ctx != nil {
					ctx = tc.ctx
				}
				_, err = s.CreateTicket(ctx, tc.req)
				tc.wantErr(t, err)
			})
		}
	})

	t.Run("Fail: Use nested transaction", func(t *testing.T) {
		fsc, err := firestoretest.InitFirestoreClient(context.Background(), "tickets")
		if err != nil {
			t.Fatalf("Failed to init firestore client: %v", err)
		}

		s, err := ticketmgr.NewTicketMgrServer(fsc)
		if err != nil {
			t.Fatalf("Failed to create server: %v", err)
		}

		fsc.RunTransaction(context.Background(), func(ctx context.Context, _ *firestore.Transaction) error {
			_, err = s.CreateTicket(ctx, baseReq)
			return nil
		})
		assert.ErrorContains(t, err, status.Errorf(codes.Internal, "failed to create ticket").Error())
	})
}

func readTicket(t testing.TB, fsc *firestore.Client, docID string) *firestorex.Ticket {
	ds, err := fsc.Doc(fmt.Sprintf("%s/%s", firestorex.CollectionNameTickets, docID)).Get(context.Background())
	if err != nil {
		t.Fatalf("Failed to read document: %v", err)
	}
	var ret firestorex.Ticket
	if err := ds.DataTo(&ret); err != nil {
		t.Fatalf("Failed to marshal document: %v", err)
	}
	return &ret
}
