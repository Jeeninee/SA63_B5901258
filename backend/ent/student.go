// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/oreo/app/ent/student"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Student holds the value of the "student" field.
	Student string `json:"student,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges StudentEdges `json:"edges"`
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// Province holds the value of the province edge.
	Province []*Province
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProvinceOrErr returns the Province value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) ProvinceOrErr() ([]*Province, error) {
	if e.loadedTypes[0] {
		return e.Province, nil
	}
	return nil, &NotLoadedError{edge: "province"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // student
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(values ...interface{}) error {
	if m, n := len(values), len(student.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field student", values[0])
	} else if value.Valid {
		s.Student = value.String
	}
	return nil
}

// QueryProvince queries the province edge of the Student.
func (s *Student) QueryProvince() *ProvinceQuery {
	return (&StudentClient{config: s.config}).QueryProvince(s)
}

// Update returns a builder for updating this Student.
// Note that, you need to call Student.Unwrap() before calling this method, if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return (&StudentClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", student=")
	builder.WriteString(s.Student)
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student

func (s Students) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}