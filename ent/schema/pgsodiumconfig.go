package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PgsodiumConfig holds the schema definition for the PgsodiumConfig entity.
type PgsodiumConfig struct {
	ent.Schema
}

// Fields of the PgsodiumConfig.
func (PgsodiumConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}).
			Comment("The ID of the project this Pgsodium config belongs to"),
		field.Bool("enabled").
			Default(false).
			Comment("Indicates if Pgsodium is enabled for the project"),
		// Additional fields as needed, such as specific Pgsodium configuration options.
	}
}

// Edges of the PgsodiumConfig.
func (PgsodiumConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("pgsodiumConfigs").
			Field("project_id").
			Unique().
			Required(),
	}
}
