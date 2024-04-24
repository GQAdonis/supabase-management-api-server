// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/networkban"
	"tribemedia.io/m/ent/project"
)

// NetworkBan is the model entity for the NetworkBan schema.
type NetworkBan struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// The ID of the project this network ban belongs to
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// The IP address or range being banned
	IPAddress string `json:"ip_address,omitempty"`
	// The reason for the ban
	Reason string `json:"reason,omitempty"`
	// The timestamp when the ban was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NetworkBanQuery when eager-loading is set.
	Edges        NetworkBanEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NetworkBanEdges holds the relations/edges for other nodes in the graph.
type NetworkBanEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NetworkBanEdges) ProjectOrErr() (*Project, error) {
	if e.Project != nil {
		return e.Project, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: project.Label}
	}
	return nil, &NotLoadedError{edge: "project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NetworkBan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case networkban.FieldIPAddress, networkban.FieldReason:
			values[i] = new(sql.NullString)
		case networkban.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case networkban.FieldID, networkban.FieldProjectID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NetworkBan fields.
func (nb *NetworkBan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case networkban.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				nb.ID = *value
			}
		case networkban.FieldProjectID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				nb.ProjectID = *value
			}
		case networkban.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				nb.IPAddress = value.String
			}
		case networkban.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				nb.Reason = value.String
			}
		case networkban.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				nb.CreatedAt = value.Time
			}
		default:
			nb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NetworkBan.
// This includes values selected through modifiers, order, etc.
func (nb *NetworkBan) Value(name string) (ent.Value, error) {
	return nb.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the NetworkBan entity.
func (nb *NetworkBan) QueryProject() *ProjectQuery {
	return NewNetworkBanClient(nb.config).QueryProject(nb)
}

// Update returns a builder for updating this NetworkBan.
// Note that you need to call NetworkBan.Unwrap() before calling this method if this NetworkBan
// was returned from a transaction, and the transaction was committed or rolled back.
func (nb *NetworkBan) Update() *NetworkBanUpdateOne {
	return NewNetworkBanClient(nb.config).UpdateOne(nb)
}

// Unwrap unwraps the NetworkBan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nb *NetworkBan) Unwrap() *NetworkBan {
	_tx, ok := nb.config.driver.(*txDriver)
	if !ok {
		panic("ent: NetworkBan is not a transactional entity")
	}
	nb.config.driver = _tx.drv
	return nb
}

// String implements the fmt.Stringer.
func (nb *NetworkBan) String() string {
	var builder strings.Builder
	builder.WriteString("NetworkBan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nb.ID))
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", nb.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("ip_address=")
	builder.WriteString(nb.IPAddress)
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(nb.Reason)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(nb.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NetworkBans is a parsable slice of NetworkBan.
type NetworkBans []*NetworkBan
