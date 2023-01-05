package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Ecomment holds the schema definition for the Ecomment entity.
type Ecomment struct {
	ent.Schema
}

// Fields of the Ecomment.
func (Ecomment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("body").
			NotEmpty(),
	}
}

// Edges of the Ecomment.
func (Ecomment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("event", Event.Type).
			Required().
			Unique(),
		edge.To("user", User.Type).
			Required().
			Unique(),
	}
}
