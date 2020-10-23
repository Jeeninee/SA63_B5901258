// Code generated by entc, DO NOT EDIT.

package subdistrict

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/oreo/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Subdistrict applies equality check predicate on the "subdistrict" field. It's identical to SubdistrictEQ.
func Subdistrict(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictEQ applies the EQ predicate on the "subdistrict" field.
func SubdistrictEQ(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictNEQ applies the NEQ predicate on the "subdistrict" field.
func SubdistrictNEQ(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictIn applies the In predicate on the "subdistrict" field.
func SubdistrictIn(vs ...string) predicate.Subdistrict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subdistrict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubdistrict), v...))
	})
}

// SubdistrictNotIn applies the NotIn predicate on the "subdistrict" field.
func SubdistrictNotIn(vs ...string) predicate.Subdistrict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subdistrict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubdistrict), v...))
	})
}

// SubdistrictGT applies the GT predicate on the "subdistrict" field.
func SubdistrictGT(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictGTE applies the GTE predicate on the "subdistrict" field.
func SubdistrictGTE(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictLT applies the LT predicate on the "subdistrict" field.
func SubdistrictLT(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictLTE applies the LTE predicate on the "subdistrict" field.
func SubdistrictLTE(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictContains applies the Contains predicate on the "subdistrict" field.
func SubdistrictContains(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictHasPrefix applies the HasPrefix predicate on the "subdistrict" field.
func SubdistrictHasPrefix(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictHasSuffix applies the HasSuffix predicate on the "subdistrict" field.
func SubdistrictHasSuffix(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictEqualFold applies the EqualFold predicate on the "subdistrict" field.
func SubdistrictEqualFold(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubdistrict), v))
	})
}

// SubdistrictContainsFold applies the ContainsFold predicate on the "subdistrict" field.
func SubdistrictContainsFold(v string) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubdistrict), v))
	})
}

// HasProvince applies the HasEdge predicate on the "province" edge.
func HasProvince() predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvinceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvinceTable, ProvinceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvinceWith applies the HasEdge predicate on the "province" edge with a given conditions (other predicates).
func HasProvinceWith(preds ...predicate.Province) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvinceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvinceTable, ProvinceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Subdistrict) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Subdistrict) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subdistrict) predicate.Subdistrict {
	return predicate.Subdistrict(func(s *sql.Selector) {
		p(s.Not())
	})
}
