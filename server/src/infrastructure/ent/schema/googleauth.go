package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GoogleAuth holds the schema definition for the GoogleAuth entity.
type GoogleAuth struct {
	ent.Schema
}

// Fields of the GoogleAuth.
func (GoogleAuth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Default(uuid.New),
		field.String("access_token").
			NotEmpty(),
		field.String("refresh_token").
			NotEmpty(),
		field.Time("expiry"),
	}
}

// Edges of the GoogleAuth.
func (GoogleAuth) Edges() []ent.Edge {
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

// Annotations of the GoogleAuth.
func (GoogleAuth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
