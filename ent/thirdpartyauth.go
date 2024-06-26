// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"tribemedia.io/m/ent/organization"
	"tribemedia.io/m/ent/thirdpartyauth"
	api "tribemedia.io/m/supabase"
)

// ThirdPartyAuth is the model entity for the ThirdPartyAuth schema.
type ThirdPartyAuth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OidcIssuerURL holds the value of the "oidc_issuer_url" field.
	OidcIssuerURL string `json:"oidc_issuer_url,omitempty"`
	// JwksURL holds the value of the "jwks_url" field.
	JwksURL string `json:"jwks_url,omitempty"`
	// CustomJwks holds the value of the "custom_jwks" field.
	CustomJwks *api.CreateThirdPartyAuthBodyCustomJwks `json:"custom_jwks,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ThirdPartyAuthQuery when eager-loading is set.
	Edges                          ThirdPartyAuthEdges `json:"edges"`
	organization_third_party_auths *int
	selectValues                   sql.SelectValues
}

// ThirdPartyAuthEdges holds the relations/edges for other nodes in the graph.
type ThirdPartyAuthEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThirdPartyAuthEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ThirdPartyAuth) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case thirdpartyauth.FieldCustomJwks:
			values[i] = new([]byte)
		case thirdpartyauth.FieldID:
			values[i] = new(sql.NullInt64)
		case thirdpartyauth.FieldOidcIssuerURL, thirdpartyauth.FieldJwksURL:
			values[i] = new(sql.NullString)
		case thirdpartyauth.ForeignKeys[0]: // organization_third_party_auths
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThirdPartyAuth fields.
func (tpa *ThirdPartyAuth) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case thirdpartyauth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tpa.ID = int(value.Int64)
		case thirdpartyauth.FieldOidcIssuerURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oidc_issuer_url", values[i])
			} else if value.Valid {
				tpa.OidcIssuerURL = value.String
			}
		case thirdpartyauth.FieldJwksURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jwks_url", values[i])
			} else if value.Valid {
				tpa.JwksURL = value.String
			}
		case thirdpartyauth.FieldCustomJwks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field custom_jwks", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tpa.CustomJwks); err != nil {
					return fmt.Errorf("unmarshal field custom_jwks: %w", err)
				}
			}
		case thirdpartyauth.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field organization_third_party_auths", value)
			} else if value.Valid {
				tpa.organization_third_party_auths = new(int)
				*tpa.organization_third_party_auths = int(value.Int64)
			}
		default:
			tpa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ThirdPartyAuth.
// This includes values selected through modifiers, order, etc.
func (tpa *ThirdPartyAuth) Value(name string) (ent.Value, error) {
	return tpa.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the ThirdPartyAuth entity.
func (tpa *ThirdPartyAuth) QueryOrganization() *OrganizationQuery {
	return NewThirdPartyAuthClient(tpa.config).QueryOrganization(tpa)
}

// Update returns a builder for updating this ThirdPartyAuth.
// Note that you need to call ThirdPartyAuth.Unwrap() before calling this method if this ThirdPartyAuth
// was returned from a transaction, and the transaction was committed or rolled back.
func (tpa *ThirdPartyAuth) Update() *ThirdPartyAuthUpdateOne {
	return NewThirdPartyAuthClient(tpa.config).UpdateOne(tpa)
}

// Unwrap unwraps the ThirdPartyAuth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tpa *ThirdPartyAuth) Unwrap() *ThirdPartyAuth {
	_tx, ok := tpa.config.driver.(*txDriver)
	if !ok {
		panic("ent: ThirdPartyAuth is not a transactional entity")
	}
	tpa.config.driver = _tx.drv
	return tpa
}

// String implements the fmt.Stringer.
func (tpa *ThirdPartyAuth) String() string {
	var builder strings.Builder
	builder.WriteString("ThirdPartyAuth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tpa.ID))
	builder.WriteString("oidc_issuer_url=")
	builder.WriteString(tpa.OidcIssuerURL)
	builder.WriteString(", ")
	builder.WriteString("jwks_url=")
	builder.WriteString(tpa.JwksURL)
	builder.WriteString(", ")
	builder.WriteString("custom_jwks=")
	builder.WriteString(fmt.Sprintf("%v", tpa.CustomJwks))
	builder.WriteByte(')')
	return builder.String()
}

// ThirdPartyAuths is a parsable slice of ThirdPartyAuth.
type ThirdPartyAuths []*ThirdPartyAuth
