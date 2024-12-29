package firestorex

import (
	"fmt"
	"time"
)

type Ticket struct {
	TicketID    string    `firestore:"TicketID" validate:"required,uuid"`
	Title       string    `firestore:"Title" validate:"required"`
	CreatedBy   string    `firestore:"CreatedBy" validate:"required,uuid"`
	CreatedAt   time.Time `firestore:"CreatedAt" validate:"required"`
	UpdatedAt   time.Time `firestore:"UpdatedAt" validate:"required"`
	Description string    `firestore:"Description" validate:"required"`
	Deadline    time.Time `firestore:"Deadline" validate:"required"`
}

func (t *Ticket) Validate() error {
	return validate.Struct(t)
}

func (t *Ticket) Path() string {
	return fmt.Sprintf("tickets/%s", t.TicketID)
}
