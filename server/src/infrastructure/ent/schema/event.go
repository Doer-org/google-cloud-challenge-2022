package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			MaxLen(100).
			NotEmpty(),
		field.String("detail").
			Optional().
			MaxLen(500),
		field.String("location").
			Optional().
			MaxLen(200),
		field.String("type").
			NotEmpty().
			MaxLen(30),
		field.String("state").
			NotEmpty().
			MaxLen(30),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin", User.Type).
			Unique(),
		edge.From("users", User.Type).
			Ref("events"),
		edge.To("comments", Comment.Type),
	}
}

// Annotations of the Event.
func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
