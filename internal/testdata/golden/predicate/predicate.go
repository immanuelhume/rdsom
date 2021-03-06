package predicate

import (
	"fmt"
)

type Predicate struct {
	Truthy bool   // true if the predicate is guaranteed to match everything
	Falsy  bool   // true if the predicate is guaranteed to match nothing
	Query  string // predicate in RediSearch query syntax
}

var TRUE = Predicate{Truthy: true, Query: "*"}
var FALSE = Predicate{Falsy: true}

func (p Predicate) And(q Predicate) Predicate {
	if p.Truthy && q.Truthy {
		return TRUE
	}
	if p.Falsy || q.Falsy {
		return FALSE
	}
	if p.Truthy {
		return q
	}
	if q.Truthy {
		return p
	}
	return Predicate{Query: fmt.Sprintf("%s %s", p.Query, q.Query)}
}

func (p *Predicate) EscapeQuery() {
	p.Query = Escape(p.Query)
}
