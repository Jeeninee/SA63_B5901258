// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/oreo/app/ent/predicate"
	"github.com/oreo/app/ent/student"
)

// StudentDelete is the builder for deleting a Student entity.
type StudentDelete struct {
	config
	hooks      []Hook
	mutation   *StudentMutation
	predicates []predicate.Student
}

// Where adds a new predicate to the delete builder.
func (sd *StudentDelete) Where(ps ...predicate.Student) *StudentDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StudentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sd.hooks) == 0 {
		affected, err = sd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sd.mutation = mutation
			affected, err = sd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sd.hooks) - 1; i >= 0; i-- {
			mut = sd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StudentDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StudentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: student.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: student.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// StudentDeleteOne is the builder for deleting a single Student entity.
type StudentDeleteOne struct {
	sd *StudentDelete
}

// Exec executes the deletion query.
func (sdo *StudentDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{student.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StudentDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}