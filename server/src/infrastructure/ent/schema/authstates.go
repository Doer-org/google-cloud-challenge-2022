package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthStates holds the schema definition for the AuthStates entity.
type AuthStates struct {
	ent.Schema
}

// Fields of the AuthStates.
func (AuthStates) Fields() []ent.Field {
	return []ent.Field{
		field.String("state").
			NotEmpty().
			MaxLen(20),
		field.String("redirect_url").
			Optional().
			MaxLen(50),
	}
}

// Edges of the AuthStates.
func (AuthStates) Edges() []ent.Edge {
	return nil
}
