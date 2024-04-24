// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"tribemedia.io/m/ent/organization"
	"tribemedia.io/m/ent/thirdpartyauth"
	api "tribemedia.io/m/supabase"
)

// ThirdPartyAuthCreate is the builder for creating a ThirdPartyAuth entity.
type ThirdPartyAuthCreate struct {
	config
	mutation *ThirdPartyAuthMutation
	hooks    []Hook
}

// SetOidcIssuerURL sets the "oidc_issuer_url" field.
func (tpac *ThirdPartyAuthCreate) SetOidcIssuerURL(s string) *ThirdPartyAuthCreate {
	tpac.mutation.SetOidcIssuerURL(s)
	return tpac
}

// SetNillableOidcIssuerURL sets the "oidc_issuer_url" field if the given value is not nil.
func (tpac *ThirdPartyAuthCreate) SetNillableOidcIssuerURL(s *string) *ThirdPartyAuthCreate {
	if s != nil {
		tpac.SetOidcIssuerURL(*s)
	}
	return tpac
}

// SetJwksURL sets the "jwks_url" field.
func (tpac *ThirdPartyAuthCreate) SetJwksURL(s string) *ThirdPartyAuthCreate {
	tpac.mutation.SetJwksURL(s)
	return tpac
}

// SetNillableJwksURL sets the "jwks_url" field if the given value is not nil.
func (tpac *ThirdPartyAuthCreate) SetNillableJwksURL(s *string) *ThirdPartyAuthCreate {
	if s != nil {
		tpac.SetJwksURL(*s)
	}
	return tpac
}

// SetCustomJwks sets the "custom_jwks" field.
func (tpac *ThirdPartyAuthCreate) SetCustomJwks(atpabcj *api.CreateThirdPartyAuthBodyCustomJwks) *ThirdPartyAuthCreate {
	tpac.mutation.SetCustomJwks(atpabcj)
	return tpac
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (tpac *ThirdPartyAuthCreate) SetOrganizationID(id int) *ThirdPartyAuthCreate {
	tpac.mutation.SetOrganizationID(id)
	return tpac
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (tpac *ThirdPartyAuthCreate) SetNillableOrganizationID(id *int) *ThirdPartyAuthCreate {
	if id != nil {
		tpac = tpac.SetOrganizationID(*id)
	}
	return tpac
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (tpac *ThirdPartyAuthCreate) SetOrganization(o *Organization) *ThirdPartyAuthCreate {
	return tpac.SetOrganizationID(o.ID)
}

// Mutation returns the ThirdPartyAuthMutation object of the builder.
func (tpac *ThirdPartyAuthCreate) Mutation() *ThirdPartyAuthMutation {
	return tpac.mutation
}

// Save creates the ThirdPartyAuth in the database.
func (tpac *ThirdPartyAuthCreate) Save(ctx context.Context) (*ThirdPartyAuth, error) {
	return withHooks(ctx, tpac.sqlSave, tpac.mutation, tpac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tpac *ThirdPartyAuthCreate) SaveX(ctx context.Context) *ThirdPartyAuth {
	v, err := tpac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpac *ThirdPartyAuthCreate) Exec(ctx context.Context) error {
	_, err := tpac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpac *ThirdPartyAuthCreate) ExecX(ctx context.Context) {
	if err := tpac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpac *ThirdPartyAuthCreate) check() error {
	return nil
}

func (tpac *ThirdPartyAuthCreate) sqlSave(ctx context.Context) (*ThirdPartyAuth, error) {
	if err := tpac.check(); err != nil {
		return nil, err
	}
	_node, _spec := tpac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tpac.mutation.id = &_node.ID
	tpac.mutation.done = true
	return _node, nil
}

func (tpac *ThirdPartyAuthCreate) createSpec() (*ThirdPartyAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &ThirdPartyAuth{config: tpac.config}
		_spec = sqlgraph.NewCreateSpec(thirdpartyauth.Table, sqlgraph.NewFieldSpec(thirdpartyauth.FieldID, field.TypeInt))
	)
	if value, ok := tpac.mutation.OidcIssuerURL(); ok {
		_spec.SetField(thirdpartyauth.FieldOidcIssuerURL, field.TypeString, value)
		_node.OidcIssuerURL = value
	}
	if value, ok := tpac.mutation.JwksURL(); ok {
		_spec.SetField(thirdpartyauth.FieldJwksURL, field.TypeString, value)
		_node.JwksURL = value
	}
	if value, ok := tpac.mutation.CustomJwks(); ok {
		_spec.SetField(thirdpartyauth.FieldCustomJwks, field.TypeJSON, value)
		_node.CustomJwks = value
	}
	if nodes := tpac.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_node.organization_third_party_auths = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThirdPartyAuthCreateBulk is the builder for creating many ThirdPartyAuth entities in bulk.
type ThirdPartyAuthCreateBulk struct {
	config
	err      error
	builders []*ThirdPartyAuthCreate
}

// Save creates the ThirdPartyAuth entities in the database.
func (tpacb *ThirdPartyAuthCreateBulk) Save(ctx context.Context) ([]*ThirdPartyAuth, error) {
	if tpacb.err != nil {
		return nil, tpacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tpacb.builders))
	nodes := make([]*ThirdPartyAuth, len(tpacb.builders))
	mutators := make([]Mutator, len(tpacb.builders))
	for i := range tpacb.builders {
		func(i int, root context.Context) {
			builder := tpacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThirdPartyAuthMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tpacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tpacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpacb *ThirdPartyAuthCreateBulk) SaveX(ctx context.Context) []*ThirdPartyAuth {
	v, err := tpacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpacb *ThirdPartyAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := tpacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpacb *ThirdPartyAuthCreateBulk) ExecX(ctx context.Context) {
	if err := tpacb.Exec(ctx); err != nil {
		panic(err)
	}
}