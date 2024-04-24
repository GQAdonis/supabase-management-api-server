// Code generated by ent, DO NOT EDIT.

package authconfig

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the authconfig type in the database.
	Label = "auth_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisableSignup holds the string denoting the disable_signup field in the database.
	FieldDisableSignup = "disable_signup"
	// FieldExternalEmailEnabled holds the string denoting the external_email_enabled field in the database.
	FieldExternalEmailEnabled = "external_email_enabled"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the authconfig in the database.
	Table = "auth_configs"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "auth_configs"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_auth_config"
)

// Columns holds all SQL columns for authconfig fields.
var Columns = []string{
	FieldID,
	FieldDisableSignup,
	FieldExternalEmailEnabled,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "auth_configs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_auth_config",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the AuthConfig queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDisableSignup orders the results by the disable_signup field.
func ByDisableSignup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisableSignup, opts...).ToFunc()
}

// ByExternalEmailEnabled orders the results by the external_email_enabled field.
func ByExternalEmailEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalEmailEnabled, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ProjectTable, ProjectColumn),
	)
}
