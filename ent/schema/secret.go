package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Secret holds the schema definition for the Secret entity.
type Secret struct {
	ent.Schema
}

// Fields of the Secret.
func (Secret) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}).
			Comment("The ID of the project this secret belongs to"),
		field.String("name").
			Comment("The name of the secret"),
		field.String("value").
			Sensitive().
			Comment("The value of the secret"),
		// Additional fields as needed.
	}
}

// Edges of the Secret.
func (Secret) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("secrets").
			Field("project_id").
			Unique().
			Required(),
	}
}
