// ent/schema/authconfig.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthConfig holds the schema definition for the AuthConfig entity.
type AuthConfig struct {
	ent.Schema
}

// Fields of the AuthConfig.
func (AuthConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bool("disable_signup"),
		field.Bool("external_email_enabled"),
		// Simplified for brevity; include other fields as necessary.
	}
}

// Edges of the AuthConfig.
func (AuthConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("auth_config").
			Unique().
			Required(),
	}
}
