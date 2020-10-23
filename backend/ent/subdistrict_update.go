// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/oreo/app/ent/predicate"
	"github.com/oreo/app/ent/province"
	"github.com/oreo/app/ent/subdistrict"
)

// SubdistrictUpdate is the builder for updating Subdistrict entities.
type SubdistrictUpdate struct {
	config
	hooks      []Hook
	mutation   *SubdistrictMutation
	predicates []predicate.Subdistrict
}

// Where adds a new predicate for the builder.
func (su *SubdistrictUpdate) Where(ps ...predicate.Subdistrict) *SubdistrictUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetSubdistrict sets the subdistrict field.
func (su *SubdistrictUpdate) SetSubdistrict(s string) *SubdistrictUpdate {
	su.mutation.SetSubdistrict(s)
	return su
}

// AddProvinceIDs adds the province edge to Province by ids.
func (su *SubdistrictUpdate) AddProvinceIDs(ids ...int) *SubdistrictUpdate {
	su.mutation.AddProvinceIDs(ids...)
	return su
}

// AddProvince adds the province edges to Province.
func (su *SubdistrictUpdate) AddProvince(p ...*Province) *SubdistrictUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddProvinceIDs(ids...)
}

// Mutation returns the SubdistrictMutation object of the builder.
func (su *SubdistrictUpdate) Mutation() *SubdistrictMutation {
	return su.mutation
}

// RemoveProvinceIDs removes the province edge to Province by ids.
func (su *SubdistrictUpdate) RemoveProvinceIDs(ids ...int) *SubdistrictUpdate {
	su.mutation.RemoveProvinceIDs(ids...)
	return su
}

// RemoveProvince removes province edges to Province.
func (su *SubdistrictUpdate) RemoveProvince(p ...*Province) *SubdistrictUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveProvinceIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SubdistrictUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Subdistrict(); ok {
		if err := subdistrict.SubdistrictValidator(v); err != nil {
			return 0, &ValidationError{Name: "subdistrict", err: fmt.Errorf("ent: validator failed for field \"subdistrict\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubdistrictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubdistrictUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubdistrictUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubdistrictUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SubdistrictUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subdistrict.Table,
			Columns: subdistrict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subdistrict.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Subdistrict(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subdistrict.FieldSubdistrict,
		})
	}
	if nodes := su.mutation.RemovedProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subdistrict.ProvinceTable,
			Columns: []string{subdistrict.ProvinceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subdistrict.ProvinceTable,
			Columns: []string{subdistrict.ProvinceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subdistrict.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SubdistrictUpdateOne is the builder for updating a single Subdistrict entity.
type SubdistrictUpdateOne struct {
	config
	hooks    []Hook
	mutation *SubdistrictMutation
}

// SetSubdistrict sets the subdistrict field.
func (suo *SubdistrictUpdateOne) SetSubdistrict(s string) *SubdistrictUpdateOne {
	suo.mutation.SetSubdistrict(s)
	return suo
}

// AddProvinceIDs adds the province edge to Province by ids.
func (suo *SubdistrictUpdateOne) AddProvinceIDs(ids ...int) *SubdistrictUpdateOne {
	suo.mutation.AddProvinceIDs(ids...)
	return suo
}

// AddProvince adds the province edges to Province.
func (suo *SubdistrictUpdateOne) AddProvince(p ...*Province) *SubdistrictUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddProvinceIDs(ids...)
}

// Mutation returns the SubdistrictMutation object of the builder.
func (suo *SubdistrictUpdateOne) Mutation() *SubdistrictMutation {
	return suo.mutation
}

// RemoveProvinceIDs removes the province edge to Province by ids.
func (suo *SubdistrictUpdateOne) RemoveProvinceIDs(ids ...int) *SubdistrictUpdateOne {
	suo.mutation.RemoveProvinceIDs(ids...)
	return suo
}

// RemoveProvince removes province edges to Province.
func (suo *SubdistrictUpdateOne) RemoveProvince(p ...*Province) *SubdistrictUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveProvinceIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SubdistrictUpdateOne) Save(ctx context.Context) (*Subdistrict, error) {
	if v, ok := suo.mutation.Subdistrict(); ok {
		if err := subdistrict.SubdistrictValidator(v); err != nil {
			return nil, &ValidationError{Name: "subdistrict", err: fmt.Errorf("ent: validator failed for field \"subdistrict\": %w", err)}
		}
	}

	var (
		err  error
		node *Subdistrict
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubdistrictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubdistrictUpdateOne) SaveX(ctx context.Context) *Subdistrict {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SubdistrictUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubdistrictUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SubdistrictUpdateOne) sqlSave(ctx context.Context) (s *Subdistrict, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subdistrict.Table,
			Columns: subdistrict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subdistrict.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Subdistrict.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Subdistrict(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subdistrict.FieldSubdistrict,
		})
	}
	if nodes := suo.mutation.RemovedProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subdistrict.ProvinceTable,
			Columns: []string{subdistrict.ProvinceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subdistrict.ProvinceTable,
			Columns: []string{subdistrict.ProvinceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Subdistrict{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subdistrict.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}