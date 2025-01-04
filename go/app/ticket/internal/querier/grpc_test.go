package querier_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	ticketquerierv1 "github.com/t-ash0410/stack-example/go/api/ticketquerier/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/modelfs"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/querier"
	"github.com/t-ash0410/stack-example/go/lib/ctxtest"
	"github.com/t-ash0410/stack-example/go/lib/firestoretest"
)

var (
	t2024_12_29_UTC = time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC)

	baseTicket = &modelfs.Ticket{
		TicketID:    "083c61da-b38d-4a8c-9c2d-f7ff466678b5",
		Title:       "Some Ticket",
		CreatedBy:   "8ea79f88-5b4b-4df6-b438-81a2ccf6b09f",
		UpdatedBy:   "8ea79f88-5b4b-4df6-b438-81a2ccf6b09f",
		Description: "Some ticket description.",
		Deadline:    t2024_12_29_UTC,
	}
)

func TestServer_QueryTickets(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error

			req *ticketquerierv1.QueryTicketsRequest

			want *ticketquerierv1.QueryTicketsResponse
		}{
			"Return tickets": {
				setupFirestore: func(c *firestore.Client) error {
					bw := c.BulkWriter(context.Background())

					ts := []*modelfs.Ticket{
						baseTicket,
						func() *modelfs.Ticket {
							r := *baseTicket
							r.TicketID = r.TicketID + "-002"
							return &r
						}(),
						func() *modelfs.Ticket {
							r := *baseTicket
							r.TicketID = "should-be-ignored"
							r.CreatedBy = "other-user"
							return &r
						}(),
					}
					for _, v := range ts {
						if _, err := bw.Create(c.Doc(v.Path()), v); err != nil {
							return err
						}
					}
					bw.End()
					return nil
				},
				req: &ticketquerierv1.QueryTicketsRequest{
					RequestedBy: baseTicket.CreatedBy,
				},
				want: &ticketquerierv1.QueryTicketsResponse{
					Tickets: []*ticketquerierv1.Ticket{
						{
							TicketId: baseTicket.TicketID,
							// CreatedAt
							// UpdatedAt
							CreatedBy:   baseTicket.CreatedBy,
							Title:       baseTicket.Title,
							Description: baseTicket.Description,
							Deadline:    timestamppb.New(baseTicket.Deadline),
						},
						{
							TicketId: baseTicket.TicketID + "-002",
							// CreatedAt
							// UpdatedAt
							CreatedBy:   baseTicket.CreatedBy,
							Title:       baseTicket.Title,
							Description: baseTicket.Description,
							Deadline:    timestamppb.New(baseTicket.Deadline),
						},
					},
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				fsc, err := firestoretest.InitFirestoreClient(context.Background(), "tickets")
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if err := tc.setupFirestore(fsc); err != nil {
					t.Fatalf("Failed to setup firestore: %v", err)
				}

				s, err := querier.NewTicketQuerierServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				res, err := s.QueryTickets(context.Background(), tc.req)
				if !assert.NoError(t, err) {
					return
				}
				if diff := cmp.Diff(tc.want, res,
					protocmp.Transform(),
					protocmp.IgnoreFields(&ticketquerierv1.Ticket{}, "created_at", "updated_at"),
				); diff != "" {
					t.Errorf("Response didn't match (-want / +got)\n%s", diff)
				}
			})
		}
	})

	t.Run("Fail", func(t *testing.T) {
		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error // optional

			ctx context.Context // optional
			req *ticketquerierv1.QueryTicketsRequest

			wantErr assert.ErrorAssertionFunc
		}{
			"Context cancelled": {
				ctx: ctxtest.CanceledContext(),
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.EqualError(t, err, context.Canceled.Error())
				},
			},
			"Internal: Failed to unmarshal": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw      = c.BulkWriter(context.Background())
						ref     = c.Doc(baseTicket.Path())
						invalid = map[string]interface{}{
							"TicketID":  "invalid",
							"Deadline":  "invalid date",
							"CreatedBy": baseTicket.CreatedBy,
						}
					)
					if _, err := bw.Create(ref, invalid); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &ticketquerierv1.QueryTicketsRequest{
					RequestedBy: baseTicket.CreatedBy,
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Error(codes.Internal, "failed to read ticket").Error())
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				fsc, err := firestoretest.InitFirestoreClient(context.Background(), "tickets")
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if tc.setupFirestore != nil {
					if err := tc.setupFirestore(fsc); err != nil {
						t.Fatalf("Failed to setup firestore: %v", err)
					}
				}

				s, err := querier.NewTicketQuerierServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				ctx := context.Background()
				if tc.ctx != nil {
					ctx = tc.ctx
				}
				_, err = s.QueryTickets(ctx, tc.req)
				tc.wantErr(t, err)
			})
		}
	})
}
