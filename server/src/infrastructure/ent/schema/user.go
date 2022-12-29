package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Min(0).
			Max(100),
		field.String("name").
			Default("unknown").
			MaxLen(20),
		field.Bool("authenticated").
			Default(false),
		field.String("gmail").
			MaxLen(50).
			Unique(),
		field.String("icon_img").
			MaxLen(200),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events",Event.Type),
	}
}
