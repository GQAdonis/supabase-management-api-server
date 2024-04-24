// Code generated by ent, DO NOT EDIT.

package typescripttype

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldLTE(FieldID, id))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldProjectID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldName, v))
}

// Definition applies equality check predicate on the "definition" field. It's identical to DefinitionEQ.
func Definition(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldDefinition, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...uuid.UUID) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNotIn(FieldProjectID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldContainsFold(FieldName, v))
}

// DefinitionEQ applies the EQ predicate on the "definition" field.
func DefinitionEQ(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEQ(FieldDefinition, v))
}

// DefinitionNEQ applies the NEQ predicate on the "definition" field.
func DefinitionNEQ(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNEQ(FieldDefinition, v))
}

// DefinitionIn applies the In predicate on the "definition" field.
func DefinitionIn(vs ...string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldIn(FieldDefinition, vs...))
}

// DefinitionNotIn applies the NotIn predicate on the "definition" field.
func DefinitionNotIn(vs ...string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldNotIn(FieldDefinition, vs...))
}

// DefinitionGT applies the GT predicate on the "definition" field.
func DefinitionGT(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldGT(FieldDefinition, v))
}

// DefinitionGTE applies the GTE predicate on the "definition" field.
func DefinitionGTE(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldGTE(FieldDefinition, v))
}

// DefinitionLT applies the LT predicate on the "definition" field.
func DefinitionLT(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldLT(FieldDefinition, v))
}

// DefinitionLTE applies the LTE predicate on the "definition" field.
func DefinitionLTE(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldLTE(FieldDefinition, v))
}

// DefinitionContains applies the Contains predicate on the "definition" field.
func DefinitionContains(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldContains(FieldDefinition, v))
}

// DefinitionHasPrefix applies the HasPrefix predicate on the "definition" field.
func DefinitionHasPrefix(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldHasPrefix(FieldDefinition, v))
}

// DefinitionHasSuffix applies the HasSuffix predicate on the "definition" field.
func DefinitionHasSuffix(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldHasSuffix(FieldDefinition, v))
}

// DefinitionEqualFold applies the EqualFold predicate on the "definition" field.
func DefinitionEqualFold(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldEqualFold(FieldDefinition, v))
}

// DefinitionContainsFold applies the ContainsFold predicate on the "definition" field.
func DefinitionContainsFold(v string) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.FieldContainsFold(FieldDefinition, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.TypeScriptType {
	return predicate.TypeScriptType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.TypeScriptType {
	return predicate.TypeScriptType(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TypeScriptType) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TypeScriptType) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TypeScriptType) predicate.TypeScriptType {
	return predicate.TypeScriptType(sql.NotPredicates(p))
}