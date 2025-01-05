package mgr

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	accountmgrv1 "github.com/t-ash0410/stack-example/go/api/accountmgr/v1"
	"github.com/t-ash0410/stack-example/go/app/account/internal/modelfs"
	"github.com/t-ash0410/stack-example/go/lib/firestorex"
)

func (s *AccountMgrServer) SlackSSO(ctx context.Context,
	req *accountmgrv1.SlackSSORequest,
) (*accountmgrv1.SlackSSOResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := &accountmgrv1.SlackSSOResponse{}

	err := s.fsc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		var (
			u   *firestorex.ResultWithMeta[modelfs.User]
			err error
			q   = s.fsc.Collection(modelfs.CollectionNameUsers).Where("Email", "==", req.Email)
		)
		for u, err = range firestorex.ReadEach[modelfs.User](tx.Documents(q)) {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to read, email = %q: %v", req.Email, err)
		}

		switch {
		// If the user account does not exist, sign up.
		case u == nil:
			u, err := newUserFromSlackSSOReq(req)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "failed to validate: %v", err)
			}
			res.UserId = u.UserID
			return tx.Create(s.fsc.Doc(u.Path()), u)

		// If the user account exists but has different slack information, update.
		case u.Data.SlackUserID != req.SlackUserId || u.Data.SlackTeamID != req.SlackTeamId:
			if err := updateUserBySlackSSOReq(u.Data, req); err != nil {
				return status.Errorf(codes.InvalidArgument, "failed to validate: %v", err)
			}
			res.UserId = u.Data.UserID
			return tx.Set(s.fsc.Doc(u.Data.Path()), u.Data)

		// If the user account exists and slack information has not changed, do
		// nothing and return the ID.
		default:
			res.UserId = u.Data.UserID
			return nil
		}
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func newUserFromSlackSSOReq(req *accountmgrv1.SlackSSORequest) (*modelfs.User, error) {
	var (
		id = uuid.NewString()
		u  = &modelfs.User{
			UserID:      id,
			CreatedBy:   id,
			UpdatedBy:   id,
			Email:       req.Email,
			Name:        req.Name,
			SlackUserID: req.SlackUserId,
			SlackTeamID: req.SlackTeamId,
		}
	)
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return u, nil
}

func updateUserBySlackSSOReq(u *modelfs.User, req *accountmgrv1.SlackSSORequest) error {
	u.UpdatedBy = u.UserID
	u.SlackUserID = req.SlackUserId
	u.SlackTeamID = req.SlackTeamId
	return u.Validate()
}
