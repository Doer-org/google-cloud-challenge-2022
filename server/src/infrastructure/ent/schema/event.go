package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(100).
			Default("unknown").
			NotEmpty(),
		field.String("detail").
			MaxLen(500),
		field.String("location").
			MaxLen(200),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("state",EState.Type).
			Unique(),
		edge.To("type",EType.Type).
			Unique(),
		edge.From("users",User.Type).
			Ref("events"),
	}
}
