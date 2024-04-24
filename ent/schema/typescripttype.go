package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TypeScriptType holds the schema definition for the TypeScriptType entity.
type TypeScriptType struct {
	ent.Schema
}

// Fields of the TypeScriptType.
func (TypeScriptType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}).
			Comment("The ID of the project this TypeScript type belongs to"),
		field.String("name").
			Comment("The name of the TypeScript type"),
		field.Text("definition").
			Comment("The TypeScript type definition"),
		// Additional fields as needed.
	}
}

// Edges of the TypeScriptType.
func (TypeScriptType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("typescriptTypes").
			Field("project_id").
			Unique().
			Required(),
	}
}
