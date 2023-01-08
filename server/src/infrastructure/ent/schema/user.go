package schema

import (
	"entgo.io/ent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/schema"
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


// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
