// Code generated by ent, DO NOT EDIT.

package customhostname

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customhostname type in the database.
	Label = "custom_hostname"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldSslStatus holds the string denoting the ssl_status field in the database.
	FieldSslStatus = "ssl_status"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the customhostname in the database.
	Table = "custom_hostnames"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "custom_hostnames"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
)

// Columns holds all SQL columns for customhostname fields.
var Columns = []string{
	FieldID,
	FieldProjectID,
	FieldHostname,
	FieldSslStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the CustomHostname queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByHostname orders the results by the hostname field.
func ByHostname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostname, opts...).ToFunc()
}

// BySslStatus orders the results by the ssl_status field.
func BySslStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSslStatus, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
