package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EState holds the schema definition for the EState entity.
type EState struct {
	ent.Schema
}

// Fields of the EState.
func (EState) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the EState.
func (EState) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Required().
			Unique().
			Ref("state"),
	}
}
