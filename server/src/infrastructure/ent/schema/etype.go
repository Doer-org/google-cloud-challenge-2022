package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EType holds the schema definition for the EType entity.
type EType struct {
	ent.Schema
}

// Fields of the EType.
func (EType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the EType.
func (EType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Unique().
			Required().
			Ref("type"),
	}
}