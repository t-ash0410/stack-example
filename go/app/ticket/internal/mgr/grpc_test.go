package mgr_test

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
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/mgr"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/modelfs"
	"github.com/t-ash0410/stack-example/go/lib/ctxtest"
	"github.com/t-ash0410/stack-example/go/lib/firestoretest"
	"github.com/t-ash0410/stack-example/go/lib/ptr"
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

func TestServer_CreateTicket(t *testing.T) {
	t.Parallel()

	var (
		baseReq = &ticketmgrv1.CreateTicketRequest{
			Title:       "Some Ticket",
			RequestedBy: "8ea79f88-5b4b-4df6-b438-81a2ccf6b09f",
			Description: "Some ticket description.",
			Deadline:    timestamppb.New(t2024_12_29_UTC.AddDate(0, 0, 10)),
		}
	)

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		cases := map[string]struct {
			req *ticketmgrv1.CreateTicketRequest

			want       *ticketmgrv1.CreateTicketResponse
			wantTicket *modelfs.Ticket
		}{
			"Create a some ticket": {
				req:  baseReq,
				want: &ticketmgrv1.CreateTicketResponse{
					// TicketId
				},
				wantTicket: &modelfs.Ticket{
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
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				s, err := mgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				res, err := s.CreateTicket(context.Background(), tc.req)
				if !assert.NoError(t, err) {
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
		t.Parallel()

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
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				s, err := mgr.NewTicketMgrServer(fsc)
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

	t.Run("Edge case: Use nested transaction", func(t *testing.T) {
		t.Parallel()

		fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
		if err != nil {
			t.Fatalf("Failed to init firestore client: %v", err)
		}

		s, err := mgr.NewTicketMgrServer(fsc)
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

func TestServer_UpdateTicket(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error

			req *ticketmgrv1.UpdateTicketRequest

			want       *ticketmgrv1.UpdateTicketResponse
			wantTicket *modelfs.Ticket
		}{
			"Update a some ticket": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw  = c.BulkWriter(context.Background())
						ref = c.Doc(baseTicket.Path())
					)
					if _, err := bw.Create(ref, baseTicket); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &ticketmgrv1.UpdateTicketRequest{
					TicketId:    baseTicket.TicketID,
					Title:       ptr.Ptr("Updated Ticket"),
					RequestedBy: "4e770fc1-0977-4ea9-911a-d67d4185817e",
					Description: ptr.Ptr("Updated ticket description."),
					Deadline:    timestamppb.New(t2024_12_29_UTC.AddDate(0, 0, 20)),
				},
				want: &ticketmgrv1.UpdateTicketResponse{},
				wantTicket: &modelfs.Ticket{
					TicketID:    baseTicket.TicketID,
					Title:       "Updated Ticket",
					CreatedBy:   baseTicket.CreatedBy,
					UpdatedBy:   "4e770fc1-0977-4ea9-911a-d67d4185817e",
					Description: "Updated ticket description.",
					Deadline:    t2024_12_29_UTC.AddDate(0, 0, 20),
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if err := tc.setupFirestore(fsc); err != nil {
					t.Fatalf("Failed to setup firestore: %v", err)
				}

				s, err := mgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				res, err := s.UpdateTicket(context.Background(), tc.req)
				if !assert.NoError(t, err) {
					return
				}
				if diff := cmp.Diff(tc.want, res, protocmp.Transform()); diff != "" {
					t.Errorf("Response didn't match (-want / +got)\n%s", diff)
					return
				}
				d := readTicket(t, fsc, tc.req.TicketId)
				if diff := cmp.Diff(tc.wantTicket, d); diff != "" {
					t.Errorf("Stored data didn't match (-want / +got)\n%s", diff)
					return
				}
			})
		}
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error

			ctx context.Context // optional
			req *ticketmgrv1.UpdateTicketRequest

			wantErr assert.ErrorAssertionFunc
		}{
			"Context cancelled": {
				ctx: ctxtest.CanceledContext(),
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.EqualError(t, err, context.Canceled.Error())
				},
			},
			"NotFound: Ticket is not found": {
				// setupFirestore // important
				req: &ticketmgrv1.UpdateTicketRequest{
					TicketId: baseTicket.TicketID,
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Error(codes.NotFound, "").Error())
				},
			},
			"FailedPrecondition: Failed to unmarshal": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw      = c.BulkWriter(context.Background())
						ref     = c.Doc(baseTicket.Path())
						invalid = map[string]interface{}{
							"Deadline": "invalid date",
						}
					)
					if _, err := bw.Create(ref, invalid); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &ticketmgrv1.UpdateTicketRequest{
					TicketId: baseTicket.TicketID,
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Error(codes.FailedPrecondition, "failed to unmarshal").Error())
				},
			},
			"InvalidArgument: Validation error": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw  = c.BulkWriter(context.Background())
						ref = c.Doc(baseTicket.Path())
					)
					if _, err := bw.Create(ref, baseTicket); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &ticketmgrv1.UpdateTicketRequest{
					TicketId:    baseTicket.TicketID,
					RequestedBy: "", // important
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Errorf(codes.InvalidArgument, "failed to validate").Error())
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if tc.setupFirestore != nil {
					if err := tc.setupFirestore(fsc); err != nil {
						t.Fatalf("Failed to setup firestore: %v", err)
					}
				}

				s, err := mgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				ctx := context.Background()
				if tc.ctx != nil {
					ctx = tc.ctx
				}
				_, err = s.UpdateTicket(ctx, tc.req)
				tc.wantErr(t, err)
			})
		}
	})
}

func TestServer_DeleteTicket(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error

			req *ticketmgrv1.DeleteTicketRequest

			want *ticketmgrv1.DeleteTicketResponse
		}{
			"Delete a some ticket": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw  = c.BulkWriter(context.Background())
						ref = c.Doc(baseTicket.Path())
					)
					if _, err := bw.Create(ref, baseTicket); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &ticketmgrv1.DeleteTicketRequest{
					TicketId: baseTicket.TicketID,
				},
				want: &ticketmgrv1.DeleteTicketResponse{},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if err := tc.setupFirestore(fsc); err != nil {
					t.Fatalf("Failed to setup firestore: %v", err)
				}

				s, err := mgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				res, err := s.DeleteTicket(context.Background(), tc.req)
				if !assert.NoError(t, err) {
					return
				}
				if diff := cmp.Diff(tc.want, res, protocmp.Transform()); diff != "" {
					t.Errorf("Response didn't match (-want / +got)\n%s", diff)
					return
				}
				_, err = fsc.Doc(fmt.Sprintf("%s/%s", modelfs.CollectionNameTickets, tc.req.TicketId)).Get(context.Background())
				if !assert.ErrorContains(t, err, status.Errorf(codes.NotFound, "").Error()) {
					return
				}
			})
		}
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		cases := map[string]struct {
			ctx context.Context // optional
			req *ticketmgrv1.DeleteTicketRequest

			wantErr assert.ErrorAssertionFunc
		}{
			"Context cancelled": {
				ctx: ctxtest.CanceledContext(),
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.EqualError(t, err, context.Canceled.Error())
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				s, err := mgr.NewTicketMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				ctx := context.Background()
				if tc.ctx != nil {
					ctx = tc.ctx
				}
				_, err = s.DeleteTicket(ctx, tc.req)
				tc.wantErr(t, err)
			})
		}
	})

	t.Run("Edge case: Use nested transaction", func(t *testing.T) {
		t.Parallel()

		fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameTickets)
		if err != nil {
			t.Fatalf("Failed to init firestore client: %v", err)
		}

		s, err := mgr.NewTicketMgrServer(fsc)
		if err != nil {
			t.Fatalf("Failed to create server: %v", err)
		}

		fsc.RunTransaction(context.Background(), func(ctx context.Context, _ *firestore.Transaction) error {
			_, err = s.DeleteTicket(ctx, &ticketmgrv1.DeleteTicketRequest{})
			return nil
		})
		assert.ErrorContains(t, err, status.Errorf(codes.Internal, "failed to delete ticket").Error())
	})
}

func readTicket(t testing.TB, fsc *firestore.Client, docID string) *modelfs.Ticket {
	ds, err := fsc.Doc(fmt.Sprintf("%s/%s", modelfs.CollectionNameTickets, docID)).Get(context.Background())
	if err != nil {
		t.Fatalf("Failed to read document: %v", err)
	}
	var ret modelfs.Ticket
	if err := ds.DataTo(&ret); err != nil {
		t.Fatalf("Failed to unmarshal document: %v", err)
	}
	return &ret
}
