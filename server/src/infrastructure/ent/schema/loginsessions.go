package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LoginSessions holds the schema definition for the LoginSessions entity.
type LoginSessions struct {
	ent.Schema
}

// Fields of the LoginSessions.
func (LoginSessions) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			MaxLen(200),
		field.UUID("user_id", uuid.UUID{}).
			Default(uuid.New),
	}
}

// Edges of the LoginSessions.
func (LoginSessions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Field("user_id").
			Required(),
	}
}
