package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			MaxLen(2000),
		field.String("redirect_url").
			Optional().
			MaxLen(300),
	}
}

// Edges of the AuthStates.
func (AuthStates) Edges() []ent.Edge {
	return nil
}

// Annotations of the AuthStates.
func (AuthStates) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
