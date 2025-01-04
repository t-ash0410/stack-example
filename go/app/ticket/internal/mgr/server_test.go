package mgr_test

import (
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/app/ticket/internal/mgr"
)

func TestNewTicketMgrServer(t *testing.T) {
	cases := map[string]struct {
		fsc *firestore.Client

		wantErr assert.ErrorAssertionFunc
	}{
		"Success": {
			fsc: &firestore.Client{},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, err, nil)
			},
		},
		"Fail: fsc is nil": {
			fsc: nil,
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "invalid arguments")
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := mgr.NewTicketMgrServer(tc.fsc)
			tc.wantErr(t, err)
		})
	}
}
