// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/oreo/app/ent/district"
	"github.com/oreo/app/ent/province"
)

// DistrictCreate is the builder for creating a District entity.
type DistrictCreate struct {
	config
	mutation *DistrictMutation
	hooks    []Hook
}

// SetDistrict sets the district field.
func (dc *DistrictCreate) SetDistrict(s string) *DistrictCreate {
	dc.mutation.SetDistrict(s)
	return dc
}

// AddProvinceIDs adds the province edge to Province by ids.
func (dc *DistrictCreate) AddProvinceIDs(ids ...int) *DistrictCreate {
	dc.mutation.AddProvinceIDs(ids...)
	return dc
}

// AddProvince adds the province edges to Province.
func (dc *DistrictCreate) AddProvince(p ...*Province) *DistrictCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddProvinceIDs(ids...)
}

// Mutation returns the DistrictMutation object of the builder.
func (dc *DistrictCreate) Mutation() *DistrictMutation {
	return dc.mutation
}

// Save creates the District in the database.
func (dc *DistrictCreate) Save(ctx context.Context) (*District, error) {
	if _, ok := dc.mutation.District(); !ok {
		return nil, &ValidationError{Name: "district", err: errors.New("ent: missing required field \"district\"")}
	}
	if v, ok := dc.mutation.District(); ok {
		if err := district.DistrictValidator(v); err != nil {
			return nil, &ValidationError{Name: "district", err: fmt.Errorf("ent: validator failed for field \"district\": %w", err)}
		}
	}
	var (
		err  error
		node *District
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DistrictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DistrictCreate) SaveX(ctx context.Context) *District {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DistrictCreate) sqlSave(ctx context.Context) (*District, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DistrictCreate) createSpec() (*District, *sqlgraph.CreateSpec) {
	var (
		d     = &District{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: district.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: district.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.District(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: district.FieldDistrict,
		})
		d.District = value
	}
	if nodes := dc.mutation.ProvinceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   district.ProvinceTable,
			Columns: []string{district.ProvinceColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
