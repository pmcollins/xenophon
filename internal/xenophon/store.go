package xenophon

import "fmt"

type store struct {
	contents    *inventory
	sellPolicy  *pricePolicy
	maxGrowRate float64
}

func newStore(c commodity) *store {
	return &store{
		contents:    &inventory{commodity: c},
		maxGrowRate: 10,
		sellPolicy: &pricePolicy{
			priceHistory: newPriceHistory(10),
			targetQty:    10,
			coeff:        1,
		},
	}
}

func (s *store) getAsk(tick int) *ask {
	if s.contents.qty < 1 {
		return nil
	}
	price := s.sellPolicy.getPriceForQty(s.contents.qty, tick)
	return &ask{
		price:     price,
		commodity: s.contents.commodity,
	}
}

func (s *store) grow() {
	pct := s.sellPolicy.getPercentBelowTarget(s.contents.qty)
	growRate := s.maxGrowRate * pct
	s.contents.qty += growRate
}

func (s *store) String() string {
	return fmt.Sprintf("{contents: %v, sellPolicy: %v}", s.contents, s.sellPolicy)
}

func (s *store) registerSale(price money, tick int) {
	s.sellPolicy.registerSale(price, tick)
}
