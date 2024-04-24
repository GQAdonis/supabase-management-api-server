package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CustomHostname holds the schema definition for the CustomHostname entity.
type CustomHostname struct {
	ent.Schema
}

// Fields of the CustomHostname.
func (CustomHostname) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}).
			Comment("The ID of the project this custom hostname belongs to"),
		field.String("hostname").
			Comment("The custom hostname"),
		field.String("ssl_status").
			Comment("The SSL status of the custom hostname"),
		// Additional fields as needed, such as verification status, etc.
	}
}

// Edges of the CustomHostname.
func (CustomHostname) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("customHostnames").
			Field("project_id").
			Unique().
			Required(),
	}
}
