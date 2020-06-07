package xenophon

import "fmt"

type pantry struct {
	contents []*sink
}

func newPantry() *pantry {
	sinks := make([]*sink, 0)
	for c := Oats; c <= Barley; c++ {
		inv := &inventory{commodity: c}
		sink := &sink{inv: inv, ratePerDay: 1}
		sinks = append(sinks, sink)
	}
	return &pantry{contents: sinks}
}

func (p *pantry) consume() {
	for _, sink := range p.contents {
		sink.consume()
	}
}

func (p *pantry) String() string {
	return fmt.Sprintf("%v", p.contents)
}
