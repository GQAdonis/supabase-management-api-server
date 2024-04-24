// Code generated by ent, DO NOT EDIT.

package pgsodiumconfig

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldLTE(FieldID, id))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldEQ(FieldProjectID, v))
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldEQ(FieldEnabled, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...uuid.UUID) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldNotIn(FieldProjectID, vs...))
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldEQ(FieldEnabled, v))
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.FieldNEQ(FieldEnabled, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PgsodiumConfig) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PgsodiumConfig) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PgsodiumConfig) predicate.PgsodiumConfig {
	return predicate.PgsodiumConfig(sql.NotPredicates(p))
}
