// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/thirdpartyauth"
)

// ThirdPartyAuthDelete is the builder for deleting a ThirdPartyAuth entity.
type ThirdPartyAuthDelete struct {
	config
	hooks    []Hook
	mutation *ThirdPartyAuthMutation
}

// Where appends a list predicates to the ThirdPartyAuthDelete builder.
func (tpad *ThirdPartyAuthDelete) Where(ps ...predicate.ThirdPartyAuth) *ThirdPartyAuthDelete {
	tpad.mutation.Where(ps...)
	return tpad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tpad *ThirdPartyAuthDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tpad.sqlExec, tpad.mutation, tpad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tpad *ThirdPartyAuthDelete) ExecX(ctx context.Context) int {
	n, err := tpad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tpad *ThirdPartyAuthDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(thirdpartyauth.Table, sqlgraph.NewFieldSpec(thirdpartyauth.FieldID, field.TypeInt))
	if ps := tpad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tpad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tpad.mutation.done = true
	return affected, err
}

// ThirdPartyAuthDeleteOne is the builder for deleting a single ThirdPartyAuth entity.
type ThirdPartyAuthDeleteOne struct {
	tpad *ThirdPartyAuthDelete
}

// Where appends a list predicates to the ThirdPartyAuthDelete builder.
func (tpado *ThirdPartyAuthDeleteOne) Where(ps ...predicate.ThirdPartyAuth) *ThirdPartyAuthDeleteOne {
	tpado.tpad.mutation.Where(ps...)
	return tpado
}

// Exec executes the deletion query.
func (tpado *ThirdPartyAuthDeleteOne) Exec(ctx context.Context) error {
	n, err := tpado.tpad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{thirdpartyauth.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tpado *ThirdPartyAuthDeleteOne) ExecX(ctx context.Context) {
	if err := tpado.Exec(ctx); err != nil {
		panic(err)
	}
}