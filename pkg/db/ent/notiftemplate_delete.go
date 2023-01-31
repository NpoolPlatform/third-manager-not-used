// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/notiftemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/predicate"
)

// NotifTemplateDelete is the builder for deleting a NotifTemplate entity.
type NotifTemplateDelete struct {
	config
	hooks    []Hook
	mutation *NotifTemplateMutation
}

// Where appends a list predicates to the NotifTemplateDelete builder.
func (ntd *NotifTemplateDelete) Where(ps ...predicate.NotifTemplate) *NotifTemplateDelete {
	ntd.mutation.Where(ps...)
	return ntd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ntd *NotifTemplateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ntd.hooks) == 0 {
		affected, err = ntd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ntd.mutation = mutation
			affected, err = ntd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ntd.hooks) - 1; i >= 0; i-- {
			if ntd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntd *NotifTemplateDelete) ExecX(ctx context.Context) int {
	n, err := ntd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ntd *NotifTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: notiftemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notiftemplate.FieldID,
			},
		},
	}
	if ps := ntd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ntd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// NotifTemplateDeleteOne is the builder for deleting a single NotifTemplate entity.
type NotifTemplateDeleteOne struct {
	ntd *NotifTemplateDelete
}

// Exec executes the deletion query.
func (ntdo *NotifTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := ntdo.ntd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notiftemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ntdo *NotifTemplateDeleteOne) ExecX(ctx context.Context) {
	ntdo.ntd.ExecX(ctx)
}