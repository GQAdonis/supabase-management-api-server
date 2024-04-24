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
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/project"
	"tribemedia.io/m/ent/typescripttype"
)

// TypeScriptTypeUpdate is the builder for updating TypeScriptType entities.
type TypeScriptTypeUpdate struct {
	config
	hooks    []Hook
	mutation *TypeScriptTypeMutation
}

// Where appends a list predicates to the TypeScriptTypeUpdate builder.
func (tstu *TypeScriptTypeUpdate) Where(ps ...predicate.TypeScriptType) *TypeScriptTypeUpdate {
	tstu.mutation.Where(ps...)
	return tstu
}

// SetProjectID sets the "project_id" field.
func (tstu *TypeScriptTypeUpdate) SetProjectID(u uuid.UUID) *TypeScriptTypeUpdate {
	tstu.mutation.SetProjectID(u)
	return tstu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (tstu *TypeScriptTypeUpdate) SetNillableProjectID(u *uuid.UUID) *TypeScriptTypeUpdate {
	if u != nil {
		tstu.SetProjectID(*u)
	}
	return tstu
}

// SetName sets the "name" field.
func (tstu *TypeScriptTypeUpdate) SetName(s string) *TypeScriptTypeUpdate {
	tstu.mutation.SetName(s)
	return tstu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tstu *TypeScriptTypeUpdate) SetNillableName(s *string) *TypeScriptTypeUpdate {
	if s != nil {
		tstu.SetName(*s)
	}
	return tstu
}

// SetDefinition sets the "definition" field.
func (tstu *TypeScriptTypeUpdate) SetDefinition(s string) *TypeScriptTypeUpdate {
	tstu.mutation.SetDefinition(s)
	return tstu
}

// SetNillableDefinition sets the "definition" field if the given value is not nil.
func (tstu *TypeScriptTypeUpdate) SetNillableDefinition(s *string) *TypeScriptTypeUpdate {
	if s != nil {
		tstu.SetDefinition(*s)
	}
	return tstu
}

// SetProject sets the "project" edge to the Project entity.
func (tstu *TypeScriptTypeUpdate) SetProject(p *Project) *TypeScriptTypeUpdate {
	return tstu.SetProjectID(p.ID)
}

// Mutation returns the TypeScriptTypeMutation object of the builder.
func (tstu *TypeScriptTypeUpdate) Mutation() *TypeScriptTypeMutation {
	return tstu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (tstu *TypeScriptTypeUpdate) ClearProject() *TypeScriptTypeUpdate {
	tstu.mutation.ClearProject()
	return tstu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tstu *TypeScriptTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tstu.sqlSave, tstu.mutation, tstu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tstu *TypeScriptTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := tstu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tstu *TypeScriptTypeUpdate) Exec(ctx context.Context) error {
	_, err := tstu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tstu *TypeScriptTypeUpdate) ExecX(ctx context.Context) {
	if err := tstu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tstu *TypeScriptTypeUpdate) check() error {
	if _, ok := tstu.mutation.ProjectID(); tstu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TypeScriptType.project"`)
	}
	return nil
}

func (tstu *TypeScriptTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tstu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(typescripttype.Table, typescripttype.Columns, sqlgraph.NewFieldSpec(typescripttype.FieldID, field.TypeUUID))
	if ps := tstu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tstu.mutation.Name(); ok {
		_spec.SetField(typescripttype.FieldName, field.TypeString, value)
	}
	if value, ok := tstu.mutation.Definition(); ok {
		_spec.SetField(typescripttype.FieldDefinition, field.TypeString, value)
	}
	if tstu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   typescripttype.ProjectTable,
			Columns: []string{typescripttype.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tstu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   typescripttype.ProjectTable,
			Columns: []string{typescripttype.ProjectColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tstu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typescripttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tstu.mutation.done = true
	return n, nil
}

// TypeScriptTypeUpdateOne is the builder for updating a single TypeScriptType entity.
type TypeScriptTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TypeScriptTypeMutation
}

// SetProjectID sets the "project_id" field.
func (tstuo *TypeScriptTypeUpdateOne) SetProjectID(u uuid.UUID) *TypeScriptTypeUpdateOne {
	tstuo.mutation.SetProjectID(u)
	return tstuo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (tstuo *TypeScriptTypeUpdateOne) SetNillableProjectID(u *uuid.UUID) *TypeScriptTypeUpdateOne {
	if u != nil {
		tstuo.SetProjectID(*u)
	}
	return tstuo
}

// SetName sets the "name" field.
func (tstuo *TypeScriptTypeUpdateOne) SetName(s string) *TypeScriptTypeUpdateOne {
	tstuo.mutation.SetName(s)
	return tstuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tstuo *TypeScriptTypeUpdateOne) SetNillableName(s *string) *TypeScriptTypeUpdateOne {
	if s != nil {
		tstuo.SetName(*s)
	}
	return tstuo
}

// SetDefinition sets the "definition" field.
func (tstuo *TypeScriptTypeUpdateOne) SetDefinition(s string) *TypeScriptTypeUpdateOne {
	tstuo.mutation.SetDefinition(s)
	return tstuo
}

// SetNillableDefinition sets the "definition" field if the given value is not nil.
func (tstuo *TypeScriptTypeUpdateOne) SetNillableDefinition(s *string) *TypeScriptTypeUpdateOne {
	if s != nil {
		tstuo.SetDefinition(*s)
	}
	return tstuo
}

// SetProject sets the "project" edge to the Project entity.
func (tstuo *TypeScriptTypeUpdateOne) SetProject(p *Project) *TypeScriptTypeUpdateOne {
	return tstuo.SetProjectID(p.ID)
}

// Mutation returns the TypeScriptTypeMutation object of the builder.
func (tstuo *TypeScriptTypeUpdateOne) Mutation() *TypeScriptTypeMutation {
	return tstuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (tstuo *TypeScriptTypeUpdateOne) ClearProject() *TypeScriptTypeUpdateOne {
	tstuo.mutation.ClearProject()
	return tstuo
}

// Where appends a list predicates to the TypeScriptTypeUpdate builder.
func (tstuo *TypeScriptTypeUpdateOne) Where(ps ...predicate.TypeScriptType) *TypeScriptTypeUpdateOne {
	tstuo.mutation.Where(ps...)
	return tstuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tstuo *TypeScriptTypeUpdateOne) Select(field string, fields ...string) *TypeScriptTypeUpdateOne {
	tstuo.fields = append([]string{field}, fields...)
	return tstuo
}

// Save executes the query and returns the updated TypeScriptType entity.
func (tstuo *TypeScriptTypeUpdateOne) Save(ctx context.Context) (*TypeScriptType, error) {
	return withHooks(ctx, tstuo.sqlSave, tstuo.mutation, tstuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tstuo *TypeScriptTypeUpdateOne) SaveX(ctx context.Context) *TypeScriptType {
	node, err := tstuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tstuo *TypeScriptTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := tstuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tstuo *TypeScriptTypeUpdateOne) ExecX(ctx context.Context) {
	if err := tstuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tstuo *TypeScriptTypeUpdateOne) check() error {
	if _, ok := tstuo.mutation.ProjectID(); tstuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TypeScriptType.project"`)
	}
	return nil
}

func (tstuo *TypeScriptTypeUpdateOne) sqlSave(ctx context.Context) (_node *TypeScriptType, err error) {
	if err := tstuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(typescripttype.Table, typescripttype.Columns, sqlgraph.NewFieldSpec(typescripttype.FieldID, field.TypeUUID))
	id, ok := tstuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TypeScriptType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tstuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, typescripttype.FieldID)
		for _, f := range fields {
			if !typescripttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != typescripttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tstuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tstuo.mutation.Name(); ok {
		_spec.SetField(typescripttype.FieldName, field.TypeString, value)
	}
	if value, ok := tstuo.mutation.Definition(); ok {
		_spec.SetField(typescripttype.FieldDefinition, field.TypeString, value)
	}
	if tstuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   typescripttype.ProjectTable,
			Columns: []string{typescripttype.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tstuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   typescripttype.ProjectTable,
			Columns: []string{typescripttype.ProjectColumn},
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
	_node = &TypeScriptType{config: tstuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tstuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typescripttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tstuo.mutation.done = true
	return _node, nil
}