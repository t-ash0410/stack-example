package firestorex_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/t-ash0410/stack-example/go/app/ticket/internal/firestorex"
)

var (
	t2024_12_29_UTC = time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC)

	baseTicket = &firestorex.Ticket{
		TicketID:    "083c61da-b38d-4a8c-9c2d-f7ff466678b5",
		Title:       "Some Ticket",
		CreatedBy:   "8ea79f88-5b4b-4df6-b438-81a2ccf6b09f",
		CreatedAt:   t2024_12_29_UTC,
		UpdatedAt:   t2024_12_29_UTC,
		Description: "Some ticket description.",
		Deadline:    t2024_12_29_UTC.AddDate(0, 0, 10),
	}
)

func TestTicket_Validate(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		if err := baseTicket.Validate(); err != nil {
			t.Errorf("Unexpected error occurred: %v", err)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		invalid := &firestorex.Ticket{
			TicketID:    "", // important
			Title:       "", // important
			CreatedBy:   "", // important
			Description: "", // important
		}
		if err := invalid.Validate(); err == nil {
			t.Errorf("Expected error, but returned nil")
		}
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
