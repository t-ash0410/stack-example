package modelfs

import (
	"fmt"
)

const CollectionNameUsers = "users"

type User struct {
	UserID      string `firestore:"UserID" validate:"required,uuid"`
	Email       string `firestore:"Email" validate:"required,email"`
	Name        string `firestore:"Name" validate:"required"`
	SlackUserID string `firestore:"SlackUserID"`
	SlackTeamID string `firestore:"SlackTeamID"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (u *User) Path() string {
	return fmt.Sprintf("%s/%s", CollectionNameUsers, u.UserID)
}
