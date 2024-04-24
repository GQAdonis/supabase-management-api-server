// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"tribemedia.io/m/ent/authconfig"
	"tribemedia.io/m/ent/predicate"
)

// AuthConfigDelete is the builder for deleting a AuthConfig entity.
type AuthConfigDelete struct {
	config
	hooks    []Hook
	mutation *AuthConfigMutation
}

// Where appends a list predicates to the AuthConfigDelete builder.
func (acd *AuthConfigDelete) Where(ps ...predicate.AuthConfig) *AuthConfigDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AuthConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AuthConfigDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AuthConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(authconfig.Table, sqlgraph.NewFieldSpec(authconfig.FieldID, field.TypeUUID))
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// AuthConfigDeleteOne is the builder for deleting a single AuthConfig entity.
type AuthConfigDeleteOne struct {
	acd *AuthConfigDelete
}

// Where appends a list predicates to the AuthConfigDelete builder.
func (acdo *AuthConfigDeleteOne) Where(ps ...predicate.AuthConfig) *AuthConfigDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *AuthConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AuthConfigDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}
