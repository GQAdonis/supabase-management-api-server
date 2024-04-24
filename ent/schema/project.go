package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		// Additional fields as needed.
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("secrets", Secret.Type),
		edge.To("typescriptTypes", TypeScriptType.Type),
		edge.To("functions", Function.Type),
		edge.To("customHostnames", CustomHostname.Type),
		edge.To("pgsodiumConfigs", PgsodiumConfig.Type),
		edge.To("networkBans", NetworkBan.Type),
		edge.To("branches", Branch.Type),
		edge.To("childBranches", Branch.Type),
		edge.To("auth_config", AuthConfig.Type).
			Unique(),
	}
}
