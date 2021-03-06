// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/oreo/app/ent/district"
)

// District is the model entity for the District schema.
type District struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// District holds the value of the "district" field.
	District string `json:"district,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DistrictQuery when eager-loading is set.
	Edges DistrictEdges `json:"edges"`
}

// DistrictEdges holds the relations/edges for other nodes in the graph.
type DistrictEdges struct {
	// Province holds the value of the province edge.
	Province []*Province
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProvinceOrErr returns the Province value or an error if the edge
// was not loaded in eager-loading.
func (e DistrictEdges) ProvinceOrErr() ([]*Province, error) {
	if e.loadedTypes[0] {
		return e.Province, nil
	}
	return nil, &NotLoadedError{edge: "province"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*District) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // district
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the District fields.
func (d *District) assignValues(values ...interface{}) error {
	if m, n := len(values), len(district.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field district", values[0])
	} else if value.Valid {
		d.District = value.String
	}
	return nil
}

// QueryProvince queries the province edge of the District.
func (d *District) QueryProvince() *ProvinceQuery {
	return (&DistrictClient{config: d.config}).QueryProvince(d)
}

// Update returns a builder for updating this District.
// Note that, you need to call District.Unwrap() before calling this method, if this District
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *District) Update() *DistrictUpdateOne {
	return (&DistrictClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *District) Unwrap() *District {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: District is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *District) String() string {
	var builder strings.Builder
	builder.WriteString("District(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", district=")
	builder.WriteString(d.District)
	builder.WriteByte(')')
	return builder.String()
}

// Districts is a parsable slice of District.
type Districts []*District

func (d Districts) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
