// Code generated by ent, DO NOT EDIT.

package project

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// HasSecrets applies the HasEdge predicate on the "secrets" edge.
func HasSecrets() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SecretsTable, SecretsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSecretsWith applies the HasEdge predicate on the "secrets" edge with a given conditions (other predicates).
func HasSecretsWith(preds ...predicate.Secret) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newSecretsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTypescriptTypes applies the HasEdge predicate on the "typescriptTypes" edge.
func HasTypescriptTypes() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TypescriptTypesTable, TypescriptTypesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTypescriptTypesWith applies the HasEdge predicate on the "typescriptTypes" edge with a given conditions (other predicates).
func HasTypescriptTypesWith(preds ...predicate.TypeScriptType) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newTypescriptTypesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFunctions applies the HasEdge predicate on the "functions" edge.
func HasFunctions() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FunctionsTable, FunctionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFunctionsWith applies the HasEdge predicate on the "functions" edge with a given conditions (other predicates).
func HasFunctionsWith(preds ...predicate.Function) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newFunctionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomHostnames applies the HasEdge predicate on the "customHostnames" edge.
func HasCustomHostnames() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomHostnamesTable, CustomHostnamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomHostnamesWith applies the HasEdge predicate on the "customHostnames" edge with a given conditions (other predicates).
func HasCustomHostnamesWith(preds ...predicate.CustomHostname) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newCustomHostnamesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPgsodiumConfigs applies the HasEdge predicate on the "pgsodiumConfigs" edge.
func HasPgsodiumConfigs() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PgsodiumConfigsTable, PgsodiumConfigsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPgsodiumConfigsWith applies the HasEdge predicate on the "pgsodiumConfigs" edge with a given conditions (other predicates).
func HasPgsodiumConfigsWith(preds ...predicate.PgsodiumConfig) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newPgsodiumConfigsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNetworkBans applies the HasEdge predicate on the "networkBans" edge.
func HasNetworkBans() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NetworkBansTable, NetworkBansColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNetworkBansWith applies the HasEdge predicate on the "networkBans" edge with a given conditions (other predicates).
func HasNetworkBansWith(preds ...predicate.NetworkBan) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newNetworkBansStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBranches applies the HasEdge predicate on the "branches" edge.
func HasBranches() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BranchesTable, BranchesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBranchesWith applies the HasEdge predicate on the "branches" edge with a given conditions (other predicates).
func HasBranchesWith(preds ...predicate.Branch) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newBranchesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildBranches applies the HasEdge predicate on the "childBranches" edge.
func HasChildBranches() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildBranchesTable, ChildBranchesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildBranchesWith applies the HasEdge predicate on the "childBranches" edge with a given conditions (other predicates).
func HasChildBranchesWith(preds ...predicate.Branch) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newChildBranchesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthConfig applies the HasEdge predicate on the "auth_config" edge.
func HasAuthConfig() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AuthConfigTable, AuthConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthConfigWith applies the HasEdge predicate on the "auth_config" edge with a given conditions (other predicates).
func HasAuthConfigWith(preds ...predicate.AuthConfig) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newAuthConfigStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(sql.NotPredicates(p))
}
