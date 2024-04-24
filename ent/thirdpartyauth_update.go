// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"tribemedia.io/m/ent/organization"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/thirdpartyauth"
	api "tribemedia.io/m/supabase"
)

// ThirdPartyAuthUpdate is the builder for updating ThirdPartyAuth entities.
type ThirdPartyAuthUpdate struct {
	config
	hooks    []Hook
	mutation *ThirdPartyAuthMutation
}

// Where appends a list predicates to the ThirdPartyAuthUpdate builder.
func (tpau *ThirdPartyAuthUpdate) Where(ps ...predicate.ThirdPartyAuth) *ThirdPartyAuthUpdate {
	tpau.mutation.Where(ps...)
	return tpau
}

// SetOidcIssuerURL sets the "oidc_issuer_url" field.
func (tpau *ThirdPartyAuthUpdate) SetOidcIssuerURL(s string) *ThirdPartyAuthUpdate {
	tpau.mutation.SetOidcIssuerURL(s)
	return tpau
}

// SetNillableOidcIssuerURL sets the "oidc_issuer_url" field if the given value is not nil.
func (tpau *ThirdPartyAuthUpdate) SetNillableOidcIssuerURL(s *string) *ThirdPartyAuthUpdate {
	if s != nil {
		tpau.SetOidcIssuerURL(*s)
	}
	return tpau
}

// ClearOidcIssuerURL clears the value of the "oidc_issuer_url" field.
func (tpau *ThirdPartyAuthUpdate) ClearOidcIssuerURL() *ThirdPartyAuthUpdate {
	tpau.mutation.ClearOidcIssuerURL()
	return tpau
}

// SetJwksURL sets the "jwks_url" field.
func (tpau *ThirdPartyAuthUpdate) SetJwksURL(s string) *ThirdPartyAuthUpdate {
	tpau.mutation.SetJwksURL(s)
	return tpau
}

// SetNillableJwksURL sets the "jwks_url" field if the given value is not nil.
func (tpau *ThirdPartyAuthUpdate) SetNillableJwksURL(s *string) *ThirdPartyAuthUpdate {
	if s != nil {
		tpau.SetJwksURL(*s)
	}
	return tpau
}

// ClearJwksURL clears the value of the "jwks_url" field.
func (tpau *ThirdPartyAuthUpdate) ClearJwksURL() *ThirdPartyAuthUpdate {
	tpau.mutation.ClearJwksURL()
	return tpau
}

// SetCustomJwks sets the "custom_jwks" field.
func (tpau *ThirdPartyAuthUpdate) SetCustomJwks(atpabcj *api.CreateThirdPartyAuthBodyCustomJwks) *ThirdPartyAuthUpdate {
	tpau.mutation.SetCustomJwks(atpabcj)
	return tpau
}

// ClearCustomJwks clears the value of the "custom_jwks" field.
func (tpau *ThirdPartyAuthUpdate) ClearCustomJwks() *ThirdPartyAuthUpdate {
	tpau.mutation.ClearCustomJwks()
	return tpau
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (tpau *ThirdPartyAuthUpdate) SetOrganizationID(id int) *ThirdPartyAuthUpdate {
	tpau.mutation.SetOrganizationID(id)
	return tpau
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (tpau *ThirdPartyAuthUpdate) SetNillableOrganizationID(id *int) *ThirdPartyAuthUpdate {
	if id != nil {
		tpau = tpau.SetOrganizationID(*id)
	}
	return tpau
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (tpau *ThirdPartyAuthUpdate) SetOrganization(o *Organization) *ThirdPartyAuthUpdate {
	return tpau.SetOrganizationID(o.ID)
}

// Mutation returns the ThirdPartyAuthMutation object of the builder.
func (tpau *ThirdPartyAuthUpdate) Mutation() *ThirdPartyAuthMutation {
	return tpau.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (tpau *ThirdPartyAuthUpdate) ClearOrganization() *ThirdPartyAuthUpdate {
	tpau.mutation.ClearOrganization()
	return tpau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tpau *ThirdPartyAuthUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tpau.sqlSave, tpau.mutation, tpau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tpau *ThirdPartyAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := tpau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tpau *ThirdPartyAuthUpdate) Exec(ctx context.Context) error {
	_, err := tpau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpau *ThirdPartyAuthUpdate) ExecX(ctx context.Context) {
	if err := tpau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tpau *ThirdPartyAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(thirdpartyauth.Table, thirdpartyauth.Columns, sqlgraph.NewFieldSpec(thirdpartyauth.FieldID, field.TypeInt))
	if ps := tpau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpau.mutation.OidcIssuerURL(); ok {
		_spec.SetField(thirdpartyauth.FieldOidcIssuerURL, field.TypeString, value)
	}
	if tpau.mutation.OidcIssuerURLCleared() {
		_spec.ClearField(thirdpartyauth.FieldOidcIssuerURL, field.TypeString)
	}
	if value, ok := tpau.mutation.JwksURL(); ok {
		_spec.SetField(thirdpartyauth.FieldJwksURL, field.TypeString, value)
	}
	if tpau.mutation.JwksURLCleared() {
		_spec.ClearField(thirdpartyauth.FieldJwksURL, field.TypeString)
	}
	if value, ok := tpau.mutation.CustomJwks(); ok {
		_spec.SetField(thirdpartyauth.FieldCustomJwks, field.TypeJSON, value)
	}
	if tpau.mutation.CustomJwksCleared() {
		_spec.ClearField(thirdpartyauth.FieldCustomJwks, field.TypeJSON)
	}
	if tpau.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thirdpartyauth.OrganizationTable,
			Columns: []string{thirdpartyauth.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpau.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thirdpartyauth.OrganizationTable,
			Columns: []string{thirdpartyauth.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tpau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thirdpartyauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tpau.mutation.done = true
	return n, nil
}

// ThirdPartyAuthUpdateOne is the builder for updating a single ThirdPartyAuth entity.
type ThirdPartyAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThirdPartyAuthMutation
}

// SetOidcIssuerURL sets the "oidc_issuer_url" field.
func (tpauo *ThirdPartyAuthUpdateOne) SetOidcIssuerURL(s string) *ThirdPartyAuthUpdateOne {
	tpauo.mutation.SetOidcIssuerURL(s)
	return tpauo
}

// SetNillableOidcIssuerURL sets the "oidc_issuer_url" field if the given value is not nil.
func (tpauo *ThirdPartyAuthUpdateOne) SetNillableOidcIssuerURL(s *string) *ThirdPartyAuthUpdateOne {
	if s != nil {
		tpauo.SetOidcIssuerURL(*s)
	}
	return tpauo
}

// ClearOidcIssuerURL clears the value of the "oidc_issuer_url" field.
func (tpauo *ThirdPartyAuthUpdateOne) ClearOidcIssuerURL() *ThirdPartyAuthUpdateOne {
	tpauo.mutation.ClearOidcIssuerURL()
	return tpauo
}

// SetJwksURL sets the "jwks_url" field.
func (tpauo *ThirdPartyAuthUpdateOne) SetJwksURL(s string) *ThirdPartyAuthUpdateOne {
	tpauo.mutation.SetJwksURL(s)
	return tpauo
}

// SetNillableJwksURL sets the "jwks_url" field if the given value is not nil.
func (tpauo *ThirdPartyAuthUpdateOne) SetNillableJwksURL(s *string) *ThirdPartyAuthUpdateOne {
	if s != nil {
		tpauo.SetJwksURL(*s)
	}
	return tpauo
}

// ClearJwksURL clears the value of the "jwks_url" field.
func (tpauo *ThirdPartyAuthUpdateOne) ClearJwksURL() *ThirdPartyAuthUpdateOne {
	tpauo.mutation.ClearJwksURL()
	return tpauo
}

// SetCustomJwks sets the "custom_jwks" field.
func (tpauo *ThirdPartyAuthUpdateOne) SetCustomJwks(atpabcj *api.CreateThirdPartyAuthBodyCustomJwks) *ThirdPartyAuthUpdateOne {
	tpauo.mutation.SetCustomJwks(atpabcj)
	return tpauo
}

// ClearCustomJwks clears the value of the "custom_jwks" field.
func (tpauo *ThirdPartyAuthUpdateOne) ClearCustomJwks() *ThirdPartyAuthUpdateOne {
	tpauo.mutation.ClearCustomJwks()
	return tpauo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (tpauo *ThirdPartyAuthUpdateOne) SetOrganizationID(id int) *ThirdPartyAuthUpdateOne {
	tpauo.mutation.SetOrganizationID(id)
	return tpauo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (tpauo *ThirdPartyAuthUpdateOne) SetNillableOrganizationID(id *int) *ThirdPartyAuthUpdateOne {
	if id != nil {
		tpauo = tpauo.SetOrganizationID(*id)
	}
	return tpauo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (tpauo *ThirdPartyAuthUpdateOne) SetOrganization(o *Organization) *ThirdPartyAuthUpdateOne {
	return tpauo.SetOrganizationID(o.ID)
}

// Mutation returns the ThirdPartyAuthMutation object of the builder.
func (tpauo *ThirdPartyAuthUpdateOne) Mutation() *ThirdPartyAuthMutation {
	return tpauo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (tpauo *ThirdPartyAuthUpdateOne) ClearOrganization() *ThirdPartyAuthUpdateOne {
	tpauo.mutation.ClearOrganization()
	return tpauo
}

// Where appends a list predicates to the ThirdPartyAuthUpdate builder.
func (tpauo *ThirdPartyAuthUpdateOne) Where(ps ...predicate.ThirdPartyAuth) *ThirdPartyAuthUpdateOne {
	tpauo.mutation.Where(ps...)
	return tpauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tpauo *ThirdPartyAuthUpdateOne) Select(field string, fields ...string) *ThirdPartyAuthUpdateOne {
	tpauo.fields = append([]string{field}, fields...)
	return tpauo
}

// Save executes the query and returns the updated ThirdPartyAuth entity.
func (tpauo *ThirdPartyAuthUpdateOne) Save(ctx context.Context) (*ThirdPartyAuth, error) {
	return withHooks(ctx, tpauo.sqlSave, tpauo.mutation, tpauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tpauo *ThirdPartyAuthUpdateOne) SaveX(ctx context.Context) *ThirdPartyAuth {
	node, err := tpauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tpauo *ThirdPartyAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := tpauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpauo *ThirdPartyAuthUpdateOne) ExecX(ctx context.Context) {
	if err := tpauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tpauo *ThirdPartyAuthUpdateOne) sqlSave(ctx context.Context) (_node *ThirdPartyAuth, err error) {
	_spec := sqlgraph.NewUpdateSpec(thirdpartyauth.Table, thirdpartyauth.Columns, sqlgraph.NewFieldSpec(thirdpartyauth.FieldID, field.TypeInt))
	id, ok := tpauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ThirdPartyAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tpauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thirdpartyauth.FieldID)
		for _, f := range fields {
			if !thirdpartyauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != thirdpartyauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tpauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpauo.mutation.OidcIssuerURL(); ok {
		_spec.SetField(thirdpartyauth.FieldOidcIssuerURL, field.TypeString, value)
	}
	if tpauo.mutation.OidcIssuerURLCleared() {
		_spec.ClearField(thirdpartyauth.FieldOidcIssuerURL, field.TypeString)
	}
	if value, ok := tpauo.mutation.JwksURL(); ok {
		_spec.SetField(thirdpartyauth.FieldJwksURL, field.TypeString, value)
	}
	if tpauo.mutation.JwksURLCleared() {
		_spec.ClearField(thirdpartyauth.FieldJwksURL, field.TypeString)
	}
	if value, ok := tpauo.mutation.CustomJwks(); ok {
		_spec.SetField(thirdpartyauth.FieldCustomJwks, field.TypeJSON, value)
	}
	if tpauo.mutation.CustomJwksCleared() {
		_spec.ClearField(thirdpartyauth.FieldCustomJwks, field.TypeJSON)
	}
	if tpauo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thirdpartyauth.OrganizationTable,
			Columns: []string{thirdpartyauth.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpauo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thirdpartyauth.OrganizationTable,
			Columns: []string{thirdpartyauth.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ThirdPartyAuth{config: tpauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tpauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thirdpartyauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tpauo.mutation.done = true
	return _node, nil
}