// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"tribemedia.io/m/ent/pgsodiumconfig"
	"tribemedia.io/m/ent/predicate"
)

// PgsodiumConfigDelete is the builder for deleting a PgsodiumConfig entity.
type PgsodiumConfigDelete struct {
	config
	hooks    []Hook
	mutation *PgsodiumConfigMutation
}

// Where appends a list predicates to the PgsodiumConfigDelete builder.
func (pcd *PgsodiumConfigDelete) Where(ps ...predicate.PgsodiumConfig) *PgsodiumConfigDelete {
	pcd.mutation.Where(ps...)
	return pcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcd *PgsodiumConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pcd.sqlExec, pcd.mutation, pcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pcd *PgsodiumConfigDelete) ExecX(ctx context.Context) int {
	n, err := pcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcd *PgsodiumConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(pgsodiumconfig.Table, sqlgraph.NewFieldSpec(pgsodiumconfig.FieldID, field.TypeUUID))
	if ps := pcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pcd.mutation.done = true
	return affected, err
}

// PgsodiumConfigDeleteOne is the builder for deleting a single PgsodiumConfig entity.
type PgsodiumConfigDeleteOne struct {
	pcd *PgsodiumConfigDelete
}

// Where appends a list predicates to the PgsodiumConfigDelete builder.
func (pcdo *PgsodiumConfigDeleteOne) Where(ps ...predicate.PgsodiumConfig) *PgsodiumConfigDeleteOne {
	pcdo.pcd.mutation.Where(ps...)
	return pcdo
}

// Exec executes the deletion query.
func (pcdo *PgsodiumConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := pcdo.pcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pgsodiumconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcdo *PgsodiumConfigDeleteOne) ExecX(ctx context.Context) {
	if err := pcdo.Exec(ctx); err != nil {
		panic(err)
	}
}