package xenophon

import (
	"fmt"
	"math"
)

type sink struct {
	inv        *inventory
	ratePerDay float64
}

func (s *sink) consume() {
	s.inv.qty -= s.ratePerDay
	s.inv.qty = math.Max(0, s.inv.qty)
}

func (s *sink) String() string {
	return fmt.Sprintf("{ratePerDay: %g, inv: %v}", s.ratePerDay, s.inv)
}
