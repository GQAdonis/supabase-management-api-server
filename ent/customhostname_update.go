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
	"tribemedia.io/m/ent/customhostname"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/project"
)

// CustomHostnameUpdate is the builder for updating CustomHostname entities.
type CustomHostnameUpdate struct {
	config
	hooks    []Hook
	mutation *CustomHostnameMutation
}

// Where appends a list predicates to the CustomHostnameUpdate builder.
func (chu *CustomHostnameUpdate) Where(ps ...predicate.CustomHostname) *CustomHostnameUpdate {
	chu.mutation.Where(ps...)
	return chu
}

// SetProjectID sets the "project_id" field.
func (chu *CustomHostnameUpdate) SetProjectID(u uuid.UUID) *CustomHostnameUpdate {
	chu.mutation.SetProjectID(u)
	return chu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (chu *CustomHostnameUpdate) SetNillableProjectID(u *uuid.UUID) *CustomHostnameUpdate {
	if u != nil {
		chu.SetProjectID(*u)
	}
	return chu
}

// SetHostname sets the "hostname" field.
func (chu *CustomHostnameUpdate) SetHostname(s string) *CustomHostnameUpdate {
	chu.mutation.SetHostname(s)
	return chu
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (chu *CustomHostnameUpdate) SetNillableHostname(s *string) *CustomHostnameUpdate {
	if s != nil {
		chu.SetHostname(*s)
	}
	return chu
}

// SetSslStatus sets the "ssl_status" field.
func (chu *CustomHostnameUpdate) SetSslStatus(s string) *CustomHostnameUpdate {
	chu.mutation.SetSslStatus(s)
	return chu
}

// SetNillableSslStatus sets the "ssl_status" field if the given value is not nil.
func (chu *CustomHostnameUpdate) SetNillableSslStatus(s *string) *CustomHostnameUpdate {
	if s != nil {
		chu.SetSslStatus(*s)
	}
	return chu
}

// SetProject sets the "project" edge to the Project entity.
func (chu *CustomHostnameUpdate) SetProject(p *Project) *CustomHostnameUpdate {
	return chu.SetProjectID(p.ID)
}

// Mutation returns the CustomHostnameMutation object of the builder.
func (chu *CustomHostnameUpdate) Mutation() *CustomHostnameMutation {
	return chu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (chu *CustomHostnameUpdate) ClearProject() *CustomHostnameUpdate {
	chu.mutation.ClearProject()
	return chu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (chu *CustomHostnameUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, chu.sqlSave, chu.mutation, chu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chu *CustomHostnameUpdate) SaveX(ctx context.Context) int {
	affected, err := chu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (chu *CustomHostnameUpdate) Exec(ctx context.Context) error {
	_, err := chu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chu *CustomHostnameUpdate) ExecX(ctx context.Context) {
	if err := chu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chu *CustomHostnameUpdate) check() error {
	if _, ok := chu.mutation.ProjectID(); chu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CustomHostname.project"`)
	}
	return nil
}

func (chu *CustomHostnameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := chu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customhostname.Table, customhostname.Columns, sqlgraph.NewFieldSpec(customhostname.FieldID, field.TypeUUID))
	if ps := chu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := chu.mutation.Hostname(); ok {
		_spec.SetField(customhostname.FieldHostname, field.TypeString, value)
	}
	if value, ok := chu.mutation.SslStatus(); ok {
		_spec.SetField(customhostname.FieldSslStatus, field.TypeString, value)
	}
	if chu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customhostname.ProjectTable,
			Columns: []string{customhostname.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := chu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customhostname.ProjectTable,
			Columns: []string{customhostname.ProjectColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, chu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customhostname.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	chu.mutation.done = true
	return n, nil
}

// CustomHostnameUpdateOne is the builder for updating a single CustomHostname entity.
type CustomHostnameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomHostnameMutation
}

// SetProjectID sets the "project_id" field.
func (chuo *CustomHostnameUpdateOne) SetProjectID(u uuid.UUID) *CustomHostnameUpdateOne {
	chuo.mutation.SetProjectID(u)
	return chuo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (chuo *CustomHostnameUpdateOne) SetNillableProjectID(u *uuid.UUID) *CustomHostnameUpdateOne {
	if u != nil {
		chuo.SetProjectID(*u)
	}
	return chuo
}

// SetHostname sets the "hostname" field.
func (chuo *CustomHostnameUpdateOne) SetHostname(s string) *CustomHostnameUpdateOne {
	chuo.mutation.SetHostname(s)
	return chuo
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (chuo *CustomHostnameUpdateOne) SetNillableHostname(s *string) *CustomHostnameUpdateOne {
	if s != nil {
		chuo.SetHostname(*s)
	}
	return chuo
}

// SetSslStatus sets the "ssl_status" field.
func (chuo *CustomHostnameUpdateOne) SetSslStatus(s string) *CustomHostnameUpdateOne {
	chuo.mutation.SetSslStatus(s)
	return chuo
}

// SetNillableSslStatus sets the "ssl_status" field if the given value is not nil.
func (chuo *CustomHostnameUpdateOne) SetNillableSslStatus(s *string) *CustomHostnameUpdateOne {
	if s != nil {
		chuo.SetSslStatus(*s)
	}
	return chuo
}

// SetProject sets the "project" edge to the Project entity.
func (chuo *CustomHostnameUpdateOne) SetProject(p *Project) *CustomHostnameUpdateOne {
	return chuo.SetProjectID(p.ID)
}

// Mutation returns the CustomHostnameMutation object of the builder.
func (chuo *CustomHostnameUpdateOne) Mutation() *CustomHostnameMutation {
	return chuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (chuo *CustomHostnameUpdateOne) ClearProject() *CustomHostnameUpdateOne {
	chuo.mutation.ClearProject()
	return chuo
}

// Where appends a list predicates to the CustomHostnameUpdate builder.
func (chuo *CustomHostnameUpdateOne) Where(ps ...predicate.CustomHostname) *CustomHostnameUpdateOne {
	chuo.mutation.Where(ps...)
	return chuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (chuo *CustomHostnameUpdateOne) Select(field string, fields ...string) *CustomHostnameUpdateOne {
	chuo.fields = append([]string{field}, fields...)
	return chuo
}

// Save executes the query and returns the updated CustomHostname entity.
func (chuo *CustomHostnameUpdateOne) Save(ctx context.Context) (*CustomHostname, error) {
	return withHooks(ctx, chuo.sqlSave, chuo.mutation, chuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chuo *CustomHostnameUpdateOne) SaveX(ctx context.Context) *CustomHostname {
	node, err := chuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (chuo *CustomHostnameUpdateOne) Exec(ctx context.Context) error {
	_, err := chuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chuo *CustomHostnameUpdateOne) ExecX(ctx context.Context) {
	if err := chuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chuo *CustomHostnameUpdateOne) check() error {
	if _, ok := chuo.mutation.ProjectID(); chuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CustomHostname.project"`)
	}
	return nil
}

func (chuo *CustomHostnameUpdateOne) sqlSave(ctx context.Context) (_node *CustomHostname, err error) {
	if err := chuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customhostname.Table, customhostname.Columns, sqlgraph.NewFieldSpec(customhostname.FieldID, field.TypeUUID))
	id, ok := chuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CustomHostname.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := chuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customhostname.FieldID)
		for _, f := range fields {
			if !customhostname.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customhostname.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := chuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := chuo.mutation.Hostname(); ok {
		_spec.SetField(customhostname.FieldHostname, field.TypeString, value)
	}
	if value, ok := chuo.mutation.SslStatus(); ok {
		_spec.SetField(customhostname.FieldSslStatus, field.TypeString, value)
	}
	if chuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customhostname.ProjectTable,
			Columns: []string{customhostname.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := chuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customhostname.ProjectTable,
			Columns: []string{customhostname.ProjectColumn},
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
	_node = &CustomHostname{config: chuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, chuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customhostname.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	chuo.mutation.done = true
	return _node, nil
}
