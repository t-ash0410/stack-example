package mgr_test

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"

	accountmgrv1 "github.com/t-ash0410/stack-example/go/api/accountmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/account/internal/mgr"
	"github.com/t-ash0410/stack-example/go/app/account/internal/modelfs"
	"github.com/t-ash0410/stack-example/go/lib/ctxtest"
	"github.com/t-ash0410/stack-example/go/lib/firestoretest"
	"github.com/t-ash0410/stack-example/go/lib/firestorex"
)

var (
	baseUser = &modelfs.User{
		UserID:      "ea9a3d7a-5793-43a8-a138-7c9b9609310c",
		CreatedBy:   "ea9a3d7a-5793-43a8-a138-7c9b9609310c",
		UpdatedBy:   "ea9a3d7a-5793-43a8-a138-7c9b9609310c",
		Email:       "john-doe@stack-example.tash0410.com",
		Name:        "John Doe",
		SlackUserID: "UXXXXXXX",
		SlackTeamID: "TXXXXXXX",
	}
)

func TestServer_SlackSSO(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error // optional

			req *accountmgrv1.SlackSSORequest

			want     *accountmgrv1.SlackSSOResponse
			wantUser *modelfs.User
		}{
			"Create a new user account and return the user id": {
				req: &accountmgrv1.SlackSSORequest{
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserId: baseUser.SlackUserID,
					SlackTeamId: baseUser.SlackTeamID,
				},
				want: &accountmgrv1.SlackSSOResponse{
					// Fill belows later
					// UserId
				},
				wantUser: &modelfs.User{
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserID: baseUser.SlackUserID,
					SlackTeamID: baseUser.SlackTeamID,
					// Fill belows later
					// UserID
					// CreatedBy
					// UpdatedBy
				},
			},
			"Update a user account and return the user id": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw  = c.BulkWriter(context.Background())
						ref = c.Doc(baseUser.Path())
					)
					if _, err := bw.Create(ref, baseUser); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &accountmgrv1.SlackSSORequest{
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserId: baseUser.SlackUserID + "-updated",
					SlackTeamId: baseUser.SlackTeamID + "-updated",
				},
				want: &accountmgrv1.SlackSSOResponse{
					UserId: baseUser.UserID,
				},
				wantUser: &modelfs.User{
					UserID:      baseUser.UserID,
					CreatedBy:   baseUser.CreatedBy,
					UpdatedBy:   baseUser.UserID,
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserID: baseUser.SlackUserID + "-updated",
					SlackTeamID: baseUser.SlackTeamID + "-updated",
				},
			},
			"Do nothing and return the user id": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw  = c.BulkWriter(context.Background())
						ref = c.Doc(baseUser.Path())
					)
					if _, err := bw.Create(ref, baseUser); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &accountmgrv1.SlackSSORequest{
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserId: baseUser.SlackUserID,
					SlackTeamId: baseUser.SlackTeamID,
				},
				want: &accountmgrv1.SlackSSOResponse{
					UserId: baseUser.UserID,
				},
				wantUser: baseUser,
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameUsers)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if tc.setupFirestore != nil {
					if err := tc.setupFirestore(fsc); err != nil {
						t.Fatalf("Failed to setup firestore: %v", err)
					}
				}

				s, err := mgr.NewAccountMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				res, err := s.SlackSSO(context.Background(), tc.req)
				if !assert.NoError(t, err) {
					return
				}

				if tc.want.UserId == "" {
					tc.want.UserId = res.UserId
				}
				if diff := cmp.Diff(tc.want, res, protocmp.Transform()); diff != "" {
					t.Errorf("Response didn't match (-want / +got)\n%s", diff)
					return
				}

				d := readUser(t, fsc, res.UserId)
				if tc.wantUser.UserID == "" {
					tc.wantUser.UserID = res.UserId
				}
				if tc.wantUser.CreatedBy == "" {
					tc.wantUser.CreatedBy = res.UserId
				}
				if tc.wantUser.UpdatedBy == "" {
					tc.wantUser.UpdatedBy = res.UserId
				}
				if diff := cmp.Diff(tc.wantUser, d); diff != "" {
					t.Errorf("Stored data didn't match (-want / +got)\n%s", diff)
				}
			})
		}
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		cases := map[string]struct {
			setupFirestore func(*firestore.Client) error // optional

			ctx context.Context // optional
			req *accountmgrv1.SlackSSORequest

			wantErr assert.ErrorAssertionFunc
		}{
			"Context cancelled": {
				ctx: ctxtest.CanceledContext(), // important
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.EqualError(t, err, context.Canceled.Error())
				},
			},
			"Internal: Failed to unmarshal": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw      = c.BulkWriter(context.Background())
						ref     = c.Doc(baseUser.Path())
						invalid = map[string]interface{}{
							"Email":     baseUser.Email,
							"CreatedBy": []string{"test"},
						}
					)
					if _, err := bw.Create(ref, invalid); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &accountmgrv1.SlackSSORequest{
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserId: baseUser.SlackUserID,
					SlackTeamId: baseUser.SlackTeamID,
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Errorf(codes.Internal, "failed to read").Error())
				},
			},
			"InvalidArgument: Failed to validate when creating user account": {
				req: &accountmgrv1.SlackSSORequest{
					Email:       baseUser.Email,
					Name:        "", // important
					SlackUserId: baseUser.SlackUserID,
					SlackTeamId: baseUser.SlackTeamID,
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Errorf(codes.InvalidArgument, "failed to validate").Error())
				},
			},
			"InvalidArgument: Failed to validate when updating user account": {
				setupFirestore: func(c *firestore.Client) error {
					var (
						bw      = c.BulkWriter(context.Background())
						ref     = c.Doc(baseUser.Path())
						invalid = map[string]interface{}{
							"Email": baseUser.Email,
						}
					)
					if _, err := bw.Create(ref, invalid); err != nil {
						return err
					}
					bw.End()
					return nil
				},
				req: &accountmgrv1.SlackSSORequest{
					Email:       baseUser.Email,
					Name:        baseUser.Name,
					SlackUserId: baseUser.SlackUserID,
					SlackTeamId: baseUser.SlackTeamID,
				},
				wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorContains(t, err, status.Errorf(codes.InvalidArgument, "failed to validate").Error())
				},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				fsc, err := firestoretest.InitFirestoreClient(context.Background(), modelfs.CollectionNameUsers)
				if err != nil {
					t.Fatalf("Failed to init firestore client: %v", err)
				}

				if tc.setupFirestore != nil {
					if err := tc.setupFirestore(fsc); err != nil {
						t.Fatalf("Failed to setup firestore: %v", err)
					}
				}

				s, err := mgr.NewAccountMgrServer(fsc)
				if err != nil {
					t.Fatalf("Failed to create server: %v", err)
				}

				ctx := context.Background()
				if tc.ctx != nil {
					ctx = tc.ctx
				}
				_, err = s.SlackSSO(ctx, tc.req)
				tc.wantErr(t, err)
			})
		}
	})
}

func readUser(t testing.TB, fsc *firestore.Client, docID string) *modelfs.User {
	ref := fsc.Doc(fmt.Sprintf("%s/%s", modelfs.CollectionNameUsers, docID))
	ret, err := firestorex.ReadOne[modelfs.User](context.Background(), ref)
	if err != nil {
		t.Fatalf("Failed to read a document: %v", err)
	}
	return ret.Data
}
