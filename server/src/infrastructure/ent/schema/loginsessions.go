package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Time("expiry"),
	}
}

// Edges of the LoginSessions.
func (LoginSessions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Field("user_id").
			Required().
			Annotations(
				entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
	}
}

// Annotations of the LoginSessions.
func (LoginSessions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
