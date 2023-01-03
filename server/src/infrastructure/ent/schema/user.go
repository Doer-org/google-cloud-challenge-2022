package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int("age").
			Optional().
			Min(0).
			Max(100),
		field.String("name").
			NotEmpty().
			MaxLen(20),
		field.Bool("authenticated").
			Default(false),
		field.String("mail").
			Optional().
			MaxLen(50),
		field.String("icon").
			NotEmpty().
			MaxLen(200),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type),
	}
}
