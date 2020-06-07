package xenophon

import "fmt"

type inventory struct {
	commodity commodity
	qty       float64
}

func (i *inventory) String() string {
	return fmt.Sprintf("{commodity: %v, qty: %g}", i.commodity, i.qty)
}
