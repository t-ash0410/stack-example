package ticketmgr_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/assert"

	ticketmgrv1 "github.com/t-ash0410/stack-example/go/api/ticketmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/ticket/internal/ticketmgr"
	"github.com/t-ash0410/stack-example/go/lib/testctx"
)

func TestServer_CreateTicket(t *testing.T) {
	t.Run("Success", func(t *testing.T) {})

	t.Run("Fail", func(t *testing.T) {
		cases := map[string]struct {
			ctx context.Context // optional
			req *ticketmgrv1.CreateTicketRequest

			wantErr assert.ErrorAssertionFunc
		}{
			"Context cancelled": {
				ctx: testctx.CanceledContext(),
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
				s, err := ticketmgr.NewTicketMgrServer(&firestore.Client{})
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
}
