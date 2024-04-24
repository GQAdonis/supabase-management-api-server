package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// NetworkBan holds the schema definition for the NetworkBan entity.
type NetworkBan struct {
	ent.Schema
}

// Fields of the NetworkBan.
func (NetworkBan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}).
			Comment("The ID of the project this network ban belongs to"),
		field.String("ip_address").
			Comment("The IP address or range being banned"),
		field.String("reason").
			Optional().
			Comment("The reason for the ban"),
		field.Time("created_at").
			Default(time.Now).
			Comment("The timestamp when the ban was created"),
		// Additional fields as needed.
	}
}

// Edges of the NetworkBan.
func (NetworkBan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("networkBans").
			Field("project_id").
			Unique().
			Required(),
	}
}
