// Code generated by entc, DO NOT EDIT.

package clubtypes

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/pattapong1/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CLUBETYPENAME applies equality check predicate on the "CLUBE_TYPE_NAME" field. It's identical to CLUBETYPENAMEEQ.
func CLUBETYPENAME(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEEQ applies the EQ predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEEQ(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMENEQ applies the NEQ predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMENEQ(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEIn applies the In predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEIn(vs ...string) predicate.ClubTypes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubTypes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCLUBETYPENAME), v...))
	})
}

// CLUBETYPENAMENotIn applies the NotIn predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMENotIn(vs ...string) predicate.ClubTypes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubTypes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCLUBETYPENAME), v...))
	})
}

// CLUBETYPENAMEGT applies the GT predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEGT(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEGTE applies the GTE predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEGTE(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMELT applies the LT predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMELT(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMELTE applies the LTE predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMELTE(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEContains applies the Contains predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEContains(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEHasPrefix applies the HasPrefix predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEHasPrefix(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEHasSuffix applies the HasSuffix predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEHasSuffix(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEEqualFold applies the EqualFold predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEEqualFold(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCLUBETYPENAME), v))
	})
}

// CLUBETYPENAMEContainsFold applies the ContainsFold predicate on the "CLUBE_TYPE_NAME" field.
func CLUBETYPENAMEContainsFold(v string) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCLUBETYPENAME), v))
	})
}

// HasClub applies the HasEdge predicate on the "club" edge.
func HasClub() predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubTable, ClubColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubWith applies the HasEdge predicate on the "club" edge with a given conditions (other predicates).
func HasClubWith(preds ...predicate.Club) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubTable, ClubColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ClubTypes) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ClubTypes) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
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
func Not(p predicate.ClubTypes) predicate.ClubTypes {
	return predicate.ClubTypes(func(s *sql.Selector) {
		p(s.Not())
	})
}
