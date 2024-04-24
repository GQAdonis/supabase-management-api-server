// Code generated by ent, DO NOT EDIT.

package branch

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the branch type in the database.
	Label = "branch"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProjectRef holds the string denoting the project_ref field in the database.
	FieldProjectRef = "project_ref"
	// FieldParentProjectRef holds the string denoting the parent_project_ref field in the database.
	FieldParentProjectRef = "parent_project_ref"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeParentProject holds the string denoting the parentproject edge name in mutations.
	EdgeParentProject = "parentProject"
	// Table holds the table name of the branch in the database.
	Table = "branches"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "branches"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_ref"
	// ParentProjectTable is the table that holds the parentProject relation/edge.
	ParentProjectTable = "branches"
	// ParentProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ParentProjectInverseTable = "projects"
	// ParentProjectColumn is the table column denoting the parentProject relation/edge.
	ParentProjectColumn = "parent_project_ref"
)

// Columns holds all SQL columns for branch fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldProjectRef,
	FieldParentProjectRef,
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

// OrderOption defines the ordering options for the Branch queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByProjectRef orders the results by the project_ref field.
func ByProjectRef(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectRef, opts...).ToFunc()
}

// ByParentProjectRef orders the results by the parent_project_ref field.
func ByParentProjectRef(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentProjectRef, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentProjectField orders the results by parentProject field.
func ByParentProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentProjectStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newParentProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParentProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentProjectTable, ParentProjectColumn),
	)
}
