// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"samsamoohooh-go-api/internal/repository/database/ent/predicate"
	"samsamoohooh-go-api/internal/repository/database/ent/topic"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicDelete is the builder for deleting a Topic entity.
type TopicDelete struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicDelete builder.
func (td *TopicDelete) Where(ps ...predicate.Topic) *TopicDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TopicDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TopicDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TopicDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(topic.Table, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TopicDeleteOne is the builder for deleting a single Topic entity.
type TopicDeleteOne struct {
	td *TopicDelete
}

// Where appends a list predicates to the TopicDelete builder.
func (tdo *TopicDeleteOne) Where(ps ...predicate.Topic) *TopicDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TopicDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{topic.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TopicDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}
