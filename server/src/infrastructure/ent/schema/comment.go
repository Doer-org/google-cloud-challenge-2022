package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("body").
			NotEmpty(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("event", Event.Type).
			Required().
			Unique().
			Annotations(
				entoas.ReadOperation(
					entoas.OperationPolicy(entoas.PolicyExclude),
				),
			),
		edge.To("user", User.Type).
			Required().
			Unique().
			Annotations(
				entoas.ReadOperation(
					entoas.OperationPolicy(entoas.PolicyExclude),
				),
			),
	}
}

// Annotations of the Comment.
func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
