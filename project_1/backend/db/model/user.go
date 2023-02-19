package model

import (
	"context"
	"github.com/google/uuid"

	"github.com/TomChv/csc-0847/project_1/backend/db"
	"github.com/TomChv/csc-0847/project_1/backend/ent"
)

// CreateUserDTO defines all fields required to create a new User.
type CreateUserDTO struct {
	StudentId      string `json:"student_id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	MailingAddress string `json:"mailing_address"`
	GPA            int    `json:"gpa"`
}

// UpdateUserDTO defines all fields required to update a new User.
type UpdateUserDTO struct {
	Firstname      string `json:"firstname,omitempty"`
	Lastname       string `json:"lastname,omitempty"`
	Email          string `json:"email,omitempty"`
	MailingAddress string `json:"mailing_address,omitempty"`
	GPA            *int   `json:"gpa,omitempty"`
}

// CreateUser inserts a new User in database.
func CreateUser(ctx context.Context, client *db.Client, data *CreateUserDTO) (*ent.User, error) {
	newUser := client.User.Create().
		SetStudentID(data.StudentId).
		SetFirstname(data.Firstname).
		SetLastname(data.Lastname).
		SetEmail(data.Email).
		SetMailingAddress(data.MailingAddress).
		SetGpa(data.GPA)

	return newUser.Save(ctx)
}

// ListUser returns all User in database.
func ListUser(ctx context.Context, client *db.Client) ([]*ent.User, error) {
	return client.User.Query().All(ctx)
}

// GetUser returns one User from database by its ID.
func GetUser(ctx context.Context, client *db.Client, id uuid.UUID) (*ent.User, error) {
	return client.User.Get(ctx, id)
}

// UpdateUser updates one User from database by its ID.
func UpdateUser(ctx context.Context, client *db.Client, id uuid.UUID, data *UpdateUserDTO) (*ent.User, error) {
	updatedUser := client.User.UpdateOneID(id)

	if data.Email != "" {
		updatedUser.SetEmail(data.Email)
	}

	if data.Firstname != "" {
		updatedUser.SetFirstname(data.Firstname)
	}

	if data.Lastname != "" {
		updatedUser.SetLastname(data.Lastname)
	}

	if data.MailingAddress != "" {
		updatedUser.SetMailingAddress(data.MailingAddress)
	}

	if data.GPA != nil {
		updatedUser.SetGpa(*data.GPA)
	}

	return updatedUser.Save(ctx)
}

// DeleteUser removes one User from database by its ID.
func DeleteUser(ctx context.Context, client *db.Client, id uuid.UUID) error {
	return client.User.DeleteOneID(id).Exec(ctx)
}
