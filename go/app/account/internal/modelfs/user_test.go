package modelfs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/app/account/internal/modelfs"
)

var (
	baseUser = &modelfs.User{
		UserID:      "b28c4f54-1a6c-4969-b3e1-e95bb794b049",
		Email:       "john-doe@stack-example.tash0410.com",
		Name:        "John Doe",
		SlackUserID: "UXXXXXXX",
		SlackTeamID: "TXXXXXXX",
	}
)

func TestUser_Validate(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, baseUser.Validate())
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		invalid := &modelfs.User{
			UserID: "", // important
			Email:  "", // important
			Name:   "", // important
		}
		assert.NotNil(t, invalid.Validate())
	})
}

func TestTicket_Path(t *testing.T) {
	t.Parallel()

	var (
		w = "users/b28c4f54-1a6c-4969-b3e1-e95bb794b049"
		g = baseUser.Path()
	)
	if diff := cmp.Diff(w, g); diff != "" {
		t.Errorf("Response did not match:(-want / +got)\n%s", diff)
	}
}
