package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Function holds the schema definition for the Function entity.
type Function struct {
	ent.Schema
}

// Fields of the Function.
func (Function) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}).
			Comment("The ID of the project this function belongs to"),
		field.String("name").
			Comment("The name of the function"),
		field.String("runtime").
			Comment("The runtime of the function"),
		field.Text("source_code").
			Sensitive().
			Comment("The source code of the function"),
		// Additional fields as needed, such as environment variables, etc.
	}
}

// Edges of the Function.
func (Function) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("functions").
			Field("project_id").
			Unique().
			Required(),
	}
}
