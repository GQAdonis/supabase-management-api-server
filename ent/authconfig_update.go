// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/authconfig"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/project"
)

// AuthConfigUpdate is the builder for updating AuthConfig entities.
type AuthConfigUpdate struct {
	config
	hooks    []Hook
	mutation *AuthConfigMutation
}

// Where appends a list predicates to the AuthConfigUpdate builder.
func (acu *AuthConfigUpdate) Where(ps ...predicate.AuthConfig) *AuthConfigUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetDisableSignup sets the "disable_signup" field.
func (acu *AuthConfigUpdate) SetDisableSignup(b bool) *AuthConfigUpdate {
	acu.mutation.SetDisableSignup(b)
	return acu
}

// SetNillableDisableSignup sets the "disable_signup" field if the given value is not nil.
func (acu *AuthConfigUpdate) SetNillableDisableSignup(b *bool) *AuthConfigUpdate {
	if b != nil {
		acu.SetDisableSignup(*b)
	}
	return acu
}

// SetExternalEmailEnabled sets the "external_email_enabled" field.
func (acu *AuthConfigUpdate) SetExternalEmailEnabled(b bool) *AuthConfigUpdate {
	acu.mutation.SetExternalEmailEnabled(b)
	return acu
}

// SetNillableExternalEmailEnabled sets the "external_email_enabled" field if the given value is not nil.
func (acu *AuthConfigUpdate) SetNillableExternalEmailEnabled(b *bool) *AuthConfigUpdate {
	if b != nil {
		acu.SetExternalEmailEnabled(*b)
	}
	return acu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (acu *AuthConfigUpdate) SetProjectID(id uuid.UUID) *AuthConfigUpdate {
	acu.mutation.SetProjectID(id)
	return acu
}

// SetProject sets the "project" edge to the Project entity.
func (acu *AuthConfigUpdate) SetProject(p *Project) *AuthConfigUpdate {
	return acu.SetProjectID(p.ID)
}

// Mutation returns the AuthConfigMutation object of the builder.
func (acu *AuthConfigUpdate) Mutation() *AuthConfigMutation {
	return acu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (acu *AuthConfigUpdate) ClearProject() *AuthConfigUpdate {
	acu.mutation.ClearProject()
	return acu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AuthConfigUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AuthConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AuthConfigUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AuthConfigUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acu *AuthConfigUpdate) check() error {
	if _, ok := acu.mutation.ProjectID(); acu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AuthConfig.project"`)
	}
	return nil
}

func (acu *AuthConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := acu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authconfig.Table, authconfig.Columns, sqlgraph.NewFieldSpec(authconfig.FieldID, field.TypeUUID))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.DisableSignup(); ok {
		_spec.SetField(authconfig.FieldDisableSignup, field.TypeBool, value)
	}
	if value, ok := acu.mutation.ExternalEmailEnabled(); ok {
		_spec.SetField(authconfig.FieldExternalEmailEnabled, field.TypeBool, value)
	}
	if acu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authconfig.ProjectTable,
			Columns: []string{authconfig.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authconfig.ProjectTable,
			Columns: []string{authconfig.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// AuthConfigUpdateOne is the builder for updating a single AuthConfig entity.
type AuthConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthConfigMutation
}

// SetDisableSignup sets the "disable_signup" field.
func (acuo *AuthConfigUpdateOne) SetDisableSignup(b bool) *AuthConfigUpdateOne {
	acuo.mutation.SetDisableSignup(b)
	return acuo
}

// SetNillableDisableSignup sets the "disable_signup" field if the given value is not nil.
func (acuo *AuthConfigUpdateOne) SetNillableDisableSignup(b *bool) *AuthConfigUpdateOne {
	if b != nil {
		acuo.SetDisableSignup(*b)
	}
	return acuo
}

// SetExternalEmailEnabled sets the "external_email_enabled" field.
func (acuo *AuthConfigUpdateOne) SetExternalEmailEnabled(b bool) *AuthConfigUpdateOne {
	acuo.mutation.SetExternalEmailEnabled(b)
	return acuo
}

// SetNillableExternalEmailEnabled sets the "external_email_enabled" field if the given value is not nil.
func (acuo *AuthConfigUpdateOne) SetNillableExternalEmailEnabled(b *bool) *AuthConfigUpdateOne {
	if b != nil {
		acuo.SetExternalEmailEnabled(*b)
	}
	return acuo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (acuo *AuthConfigUpdateOne) SetProjectID(id uuid.UUID) *AuthConfigUpdateOne {
	acuo.mutation.SetProjectID(id)
	return acuo
}

// SetProject sets the "project" edge to the Project entity.
func (acuo *AuthConfigUpdateOne) SetProject(p *Project) *AuthConfigUpdateOne {
	return acuo.SetProjectID(p.ID)
}

// Mutation returns the AuthConfigMutation object of the builder.
func (acuo *AuthConfigUpdateOne) Mutation() *AuthConfigMutation {
	return acuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (acuo *AuthConfigUpdateOne) ClearProject() *AuthConfigUpdateOne {
	acuo.mutation.ClearProject()
	return acuo
}

// Where appends a list predicates to the AuthConfigUpdate builder.
func (acuo *AuthConfigUpdateOne) Where(ps ...predicate.AuthConfig) *AuthConfigUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AuthConfigUpdateOne) Select(field string, fields ...string) *AuthConfigUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AuthConfig entity.
func (acuo *AuthConfigUpdateOne) Save(ctx context.Context) (*AuthConfig, error) {
	return withHooks(ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AuthConfigUpdateOne) SaveX(ctx context.Context) *AuthConfig {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AuthConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AuthConfigUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acuo *AuthConfigUpdateOne) check() error {
	if _, ok := acuo.mutation.ProjectID(); acuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AuthConfig.project"`)
	}
	return nil
}

func (acuo *AuthConfigUpdateOne) sqlSave(ctx context.Context) (_node *AuthConfig, err error) {
	if err := acuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authconfig.Table, authconfig.Columns, sqlgraph.NewFieldSpec(authconfig.FieldID, field.TypeUUID))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authconfig.FieldID)
		for _, f := range fields {
			if !authconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.DisableSignup(); ok {
		_spec.SetField(authconfig.FieldDisableSignup, field.TypeBool, value)
	}
	if value, ok := acuo.mutation.ExternalEmailEnabled(); ok {
		_spec.SetField(authconfig.FieldExternalEmailEnabled, field.TypeBool, value)
	}
	if acuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authconfig.ProjectTable,
			Columns: []string{authconfig.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authconfig.ProjectTable,
			Columns: []string{authconfig.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthConfig{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}
