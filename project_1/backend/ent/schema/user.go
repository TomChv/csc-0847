package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("student_id").Unique(),
		field.String("firstname"),
		field.String("lastname"),
		field.String("email").
			Unique().
			Match(regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)),
		field.String("mailing_address"),
		field.Int("gpa").
			Min(0).
			Max(4),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
