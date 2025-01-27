package modelfs_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/app/ticket/internal/modelfs"
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

func TestTicket_Validate(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, baseTicket.Validate())
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		invalid := &modelfs.Ticket{
			TicketID:    "", // important
			Title:       "", // important
			CreatedBy:   "", // important
			Description: "", // important
		}
		assert.NotNil(t, invalid.Validate())
	})
}

func TestTicket_Path(t *testing.T) {
	t.Parallel()

	var (
		w = "tickets/083c61da-b38d-4a8c-9c2d-f7ff466678b5"
		g = baseTicket.Path()
	)
	if diff := cmp.Diff(w, g); diff != "" {
		t.Errorf("Response did not match:(-want / +got)\n%s", diff)
	}
}
