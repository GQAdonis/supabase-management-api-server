// Code generated by ent, DO NOT EDIT.

package function

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldLTE(FieldID, id))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldProjectID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldName, v))
}

// Runtime applies equality check predicate on the "runtime" field. It's identical to RuntimeEQ.
func Runtime(v string) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldRuntime, v))
}

// SourceCode applies equality check predicate on the "source_code" field. It's identical to SourceCodeEQ.
func SourceCode(v string) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldSourceCode, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...uuid.UUID) predicate.Function {
	return predicate.Function(sql.FieldNotIn(FieldProjectID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Function {
	return predicate.Function(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Function {
	return predicate.Function(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Function {
	return predicate.Function(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Function {
	return predicate.Function(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Function {
	return predicate.Function(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Function {
	return predicate.Function(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Function {
	return predicate.Function(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Function {
	return predicate.Function(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Function {
	return predicate.Function(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Function {
	return predicate.Function(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Function {
	return predicate.Function(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Function {
	return predicate.Function(sql.FieldContainsFold(FieldName, v))
}

// RuntimeEQ applies the EQ predicate on the "runtime" field.
func RuntimeEQ(v string) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldRuntime, v))
}

// RuntimeNEQ applies the NEQ predicate on the "runtime" field.
func RuntimeNEQ(v string) predicate.Function {
	return predicate.Function(sql.FieldNEQ(FieldRuntime, v))
}

// RuntimeIn applies the In predicate on the "runtime" field.
func RuntimeIn(vs ...string) predicate.Function {
	return predicate.Function(sql.FieldIn(FieldRuntime, vs...))
}

// RuntimeNotIn applies the NotIn predicate on the "runtime" field.
func RuntimeNotIn(vs ...string) predicate.Function {
	return predicate.Function(sql.FieldNotIn(FieldRuntime, vs...))
}

// RuntimeGT applies the GT predicate on the "runtime" field.
func RuntimeGT(v string) predicate.Function {
	return predicate.Function(sql.FieldGT(FieldRuntime, v))
}

// RuntimeGTE applies the GTE predicate on the "runtime" field.
func RuntimeGTE(v string) predicate.Function {
	return predicate.Function(sql.FieldGTE(FieldRuntime, v))
}

// RuntimeLT applies the LT predicate on the "runtime" field.
func RuntimeLT(v string) predicate.Function {
	return predicate.Function(sql.FieldLT(FieldRuntime, v))
}

// RuntimeLTE applies the LTE predicate on the "runtime" field.
func RuntimeLTE(v string) predicate.Function {
	return predicate.Function(sql.FieldLTE(FieldRuntime, v))
}

// RuntimeContains applies the Contains predicate on the "runtime" field.
func RuntimeContains(v string) predicate.Function {
	return predicate.Function(sql.FieldContains(FieldRuntime, v))
}

// RuntimeHasPrefix applies the HasPrefix predicate on the "runtime" field.
func RuntimeHasPrefix(v string) predicate.Function {
	return predicate.Function(sql.FieldHasPrefix(FieldRuntime, v))
}

// RuntimeHasSuffix applies the HasSuffix predicate on the "runtime" field.
func RuntimeHasSuffix(v string) predicate.Function {
	return predicate.Function(sql.FieldHasSuffix(FieldRuntime, v))
}

// RuntimeEqualFold applies the EqualFold predicate on the "runtime" field.
func RuntimeEqualFold(v string) predicate.Function {
	return predicate.Function(sql.FieldEqualFold(FieldRuntime, v))
}

// RuntimeContainsFold applies the ContainsFold predicate on the "runtime" field.
func RuntimeContainsFold(v string) predicate.Function {
	return predicate.Function(sql.FieldContainsFold(FieldRuntime, v))
}

// SourceCodeEQ applies the EQ predicate on the "source_code" field.
func SourceCodeEQ(v string) predicate.Function {
	return predicate.Function(sql.FieldEQ(FieldSourceCode, v))
}

// SourceCodeNEQ applies the NEQ predicate on the "source_code" field.
func SourceCodeNEQ(v string) predicate.Function {
	return predicate.Function(sql.FieldNEQ(FieldSourceCode, v))
}

// SourceCodeIn applies the In predicate on the "source_code" field.
func SourceCodeIn(vs ...string) predicate.Function {
	return predicate.Function(sql.FieldIn(FieldSourceCode, vs...))
}

// SourceCodeNotIn applies the NotIn predicate on the "source_code" field.
func SourceCodeNotIn(vs ...string) predicate.Function {
	return predicate.Function(sql.FieldNotIn(FieldSourceCode, vs...))
}

// SourceCodeGT applies the GT predicate on the "source_code" field.
func SourceCodeGT(v string) predicate.Function {
	return predicate.Function(sql.FieldGT(FieldSourceCode, v))
}

// SourceCodeGTE applies the GTE predicate on the "source_code" field.
func SourceCodeGTE(v string) predicate.Function {
	return predicate.Function(sql.FieldGTE(FieldSourceCode, v))
}

// SourceCodeLT applies the LT predicate on the "source_code" field.
func SourceCodeLT(v string) predicate.Function {
	return predicate.Function(sql.FieldLT(FieldSourceCode, v))
}

// SourceCodeLTE applies the LTE predicate on the "source_code" field.
func SourceCodeLTE(v string) predicate.Function {
	return predicate.Function(sql.FieldLTE(FieldSourceCode, v))
}

// SourceCodeContains applies the Contains predicate on the "source_code" field.
func SourceCodeContains(v string) predicate.Function {
	return predicate.Function(sql.FieldContains(FieldSourceCode, v))
}

// SourceCodeHasPrefix applies the HasPrefix predicate on the "source_code" field.
func SourceCodeHasPrefix(v string) predicate.Function {
	return predicate.Function(sql.FieldHasPrefix(FieldSourceCode, v))
}

// SourceCodeHasSuffix applies the HasSuffix predicate on the "source_code" field.
func SourceCodeHasSuffix(v string) predicate.Function {
	return predicate.Function(sql.FieldHasSuffix(FieldSourceCode, v))
}

// SourceCodeEqualFold applies the EqualFold predicate on the "source_code" field.
func SourceCodeEqualFold(v string) predicate.Function {
	return predicate.Function(sql.FieldEqualFold(FieldSourceCode, v))
}

// SourceCodeContainsFold applies the ContainsFold predicate on the "source_code" field.
func SourceCodeContainsFold(v string) predicate.Function {
	return predicate.Function(sql.FieldContainsFold(FieldSourceCode, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Function {
	return predicate.Function(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Function {
	return predicate.Function(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Function) predicate.Function {
	return predicate.Function(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Function) predicate.Function {
	return predicate.Function(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Function) predicate.Function {
	return predicate.Function(sql.NotPredicates(p))
}