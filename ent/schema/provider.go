// ent/schema/provider.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Provider holds the schema definition for the Provider entity.
type Provider struct {
	ent.Schema
}

// Fields of the Provider.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("metadata_xml").Optional(),
		field.String("metadata_url").Optional(),
		field.JSON("domains", []string{}),
		field.JSON("attribute_mapping", map[string]string{}).Optional(),
	}
}

// Edges of the Provider.
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("providers").
			Unique(),
	}
}
