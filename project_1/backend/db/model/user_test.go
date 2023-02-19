package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/TomChv/csc-0847/project_1/backend/db"
	"github.com/TomChv/csc-0847/project_1/backend/ent"
)

func compareUser(t *testing.T, expected *ent.User, result *ent.User) {
	assert.Equal(t, expected.StudentID, result.StudentID)
	assert.Equal(t, expected.Email, result.Email)
	assert.Equal(t, expected.Firstname, result.Firstname)
	assert.Equal(t, expected.Lastname, result.Lastname)
	assert.Equal(t, expected.MailingAddress, result.MailingAddress)
	assert.Equal(t, expected.Gpa, result.Gpa)
}

func TestCreateUser(t *testing.T) {
	t.Setenv("UNIT_TEST", "go")
	t.Setenv("DB_PROVIDER", "sqlite")

	c, err := db.New()
	assert.NoError(t, err)

	defer c.Close()

	type TestCase struct {
		input    *CreateUserDTO
		expected *ent.User
	}

	testCases := map[string]TestCase{
		"shall create user": {
			input: &CreateUserDTO{
				StudentId:      "030493542",
				Email:          "030493542@sfsu.edu",
				Firstname:      "John",
				Lastname:       "Doe",
				GPA:            3,
				MailingAddress: "No where",
			},
			expected: &ent.User{
				StudentID:      "030493542",
				Email:          "030493542@sfsu.edu",
				Firstname:      "John",
				Lastname:       "Doe",
				Gpa:            3,
				MailingAddress: "No where",
			},
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			result, err := CreateUser(ctx, c, tt.input)
			assert.NoError(t, err)
			compareUser(t, tt.expected, result)
		})
	}
}
