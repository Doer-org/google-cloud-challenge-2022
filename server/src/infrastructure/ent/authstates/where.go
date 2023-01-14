// Code generated by ent, DO NOT EDIT.

package authstates

import (
	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// RedirectURL applies equality check predicate on the "redirect_url" field. It's identical to RedirectURLEQ.
func RedirectURL(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.AuthStates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.AuthStates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// RedirectURLEQ applies the EQ predicate on the "redirect_url" field.
func RedirectURLEQ(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLNEQ applies the NEQ predicate on the "redirect_url" field.
func RedirectURLNEQ(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLIn applies the In predicate on the "redirect_url" field.
func RedirectURLIn(vs ...string) predicate.AuthStates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLNotIn applies the NotIn predicate on the "redirect_url" field.
func RedirectURLNotIn(vs ...string) predicate.AuthStates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLGT applies the GT predicate on the "redirect_url" field.
func RedirectURLGT(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLGTE applies the GTE predicate on the "redirect_url" field.
func RedirectURLGTE(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLT applies the LT predicate on the "redirect_url" field.
func RedirectURLLT(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLTE applies the LTE predicate on the "redirect_url" field.
func RedirectURLLTE(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContains applies the Contains predicate on the "redirect_url" field.
func RedirectURLContains(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasPrefix applies the HasPrefix predicate on the "redirect_url" field.
func RedirectURLHasPrefix(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasSuffix applies the HasSuffix predicate on the "redirect_url" field.
func RedirectURLHasSuffix(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLIsNil applies the IsNil predicate on the "redirect_url" field.
func RedirectURLIsNil() predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRedirectURL)))
	})
}

// RedirectURLNotNil applies the NotNil predicate on the "redirect_url" field.
func RedirectURLNotNil() predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRedirectURL)))
	})
}

// RedirectURLEqualFold applies the EqualFold predicate on the "redirect_url" field.
func RedirectURLEqualFold(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContainsFold applies the ContainsFold predicate on the "redirect_url" field.
func RedirectURLContainsFold(v string) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRedirectURL), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AuthStates) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AuthStates) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
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
func Not(p predicate.AuthStates) predicate.AuthStates {
	return predicate.AuthStates(func(s *sql.Selector) {
		p(s.Not())
	})
}