// Code generated by ent, DO NOT EDIT.

package networkban

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLTE(FieldID, id))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldProjectID, v))
}

// IPAddress applies equality check predicate on the "ip_address" field. It's identical to IPAddressEQ.
func IPAddress(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldIPAddress, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldReason, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldCreatedAt, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...uuid.UUID) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNotIn(FieldProjectID, vs...))
}

// IPAddressEQ applies the EQ predicate on the "ip_address" field.
func IPAddressEQ(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldIPAddress, v))
}

// IPAddressNEQ applies the NEQ predicate on the "ip_address" field.
func IPAddressNEQ(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNEQ(FieldIPAddress, v))
}

// IPAddressIn applies the In predicate on the "ip_address" field.
func IPAddressIn(vs ...string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldIn(FieldIPAddress, vs...))
}

// IPAddressNotIn applies the NotIn predicate on the "ip_address" field.
func IPAddressNotIn(vs ...string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNotIn(FieldIPAddress, vs...))
}

// IPAddressGT applies the GT predicate on the "ip_address" field.
func IPAddressGT(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGT(FieldIPAddress, v))
}

// IPAddressGTE applies the GTE predicate on the "ip_address" field.
func IPAddressGTE(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGTE(FieldIPAddress, v))
}

// IPAddressLT applies the LT predicate on the "ip_address" field.
func IPAddressLT(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLT(FieldIPAddress, v))
}

// IPAddressLTE applies the LTE predicate on the "ip_address" field.
func IPAddressLTE(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLTE(FieldIPAddress, v))
}

// IPAddressContains applies the Contains predicate on the "ip_address" field.
func IPAddressContains(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldContains(FieldIPAddress, v))
}

// IPAddressHasPrefix applies the HasPrefix predicate on the "ip_address" field.
func IPAddressHasPrefix(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldHasPrefix(FieldIPAddress, v))
}

// IPAddressHasSuffix applies the HasSuffix predicate on the "ip_address" field.
func IPAddressHasSuffix(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldHasSuffix(FieldIPAddress, v))
}

// IPAddressEqualFold applies the EqualFold predicate on the "ip_address" field.
func IPAddressEqualFold(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEqualFold(FieldIPAddress, v))
}

// IPAddressContainsFold applies the ContainsFold predicate on the "ip_address" field.
func IPAddressContainsFold(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldContainsFold(FieldIPAddress, v))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonIsNil applies the IsNil predicate on the "reason" field.
func ReasonIsNil() predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldIsNull(FieldReason))
}

// ReasonNotNil applies the NotNil predicate on the "reason" field.
func ReasonNotNil() predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNotNull(FieldReason))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldContainsFold(FieldReason, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NetworkBan {
	return predicate.NetworkBan(sql.FieldLTE(FieldCreatedAt, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.NetworkBan {
	return predicate.NetworkBan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.NetworkBan {
	return predicate.NetworkBan(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NetworkBan) predicate.NetworkBan {
	return predicate.NetworkBan(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NetworkBan) predicate.NetworkBan {
	return predicate.NetworkBan(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NetworkBan) predicate.NetworkBan {
	return predicate.NetworkBan(sql.NotPredicates(p))
}
