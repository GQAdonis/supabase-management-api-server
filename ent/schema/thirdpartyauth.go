// ent/schema/thirdpartyauth.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	supabase "tribemedia.io/m/supabase"
)

// ThirdPartyAuth holds the schema definition for the ThirdPartyAuth entity.
type ThirdPartyAuth struct {
	ent.Schema
}

// Fields of the ThirdPartyAuth.
func (ThirdPartyAuth) Fields() []ent.Field {
	return []ent.Field{
		field.String("oidc_issuer_url").Optional(),
		field.String("jwks_url").Optional(),
		field.JSON("custom_jwks", &supabase.CreateThirdPartyAuthBodyCustomJwks{}).Optional(),
	}
}

// Edges of the ThirdPartyAuth.
func (ThirdPartyAuth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("third_party_auths").
			Unique(),
	}
}
