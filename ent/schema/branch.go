// ent/schema/branch.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Branch holds the schema definition for the Branch entity.
type Branch struct {
	ent.Schema
}

// Fields of the Branch.
func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.UUID("project_ref", uuid.UUID{}).Optional(),
		field.UUID("parent_project_ref", uuid.UUID{}).Optional(),
		// Other fields...
	}
}

// Edges of the Branch.
func (Branch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("branches").
			Unique().
			Field("project_ref"),
		edge.From("parentProject", Project.Type).
			Ref("childBranches").
			Unique().
			Field("parent_project_ref"),
	}
}
