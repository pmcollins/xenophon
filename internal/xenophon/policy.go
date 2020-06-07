package xenophon

import (
	"fmt"
)

type pricePolicyMap map[commodity]*pricePolicy

type buyPolicyCatalog struct {
	policies pricePolicyMap
}

func newBuyPolicyCatalog() *buyPolicyCatalog {
	p := &pricePolicy{
		priceHistory: newPriceHistory(10),
		targetQty:    10,
		coeff:        1,
	}
	return &buyPolicyCatalog{policies: pricePolicyMap{
		Oats:   p,
		Peas:   p,
		Beans:  p,
		Barley: p,
	}}
}

func (p *buyPolicyCatalog) getPrice(inv *inventory, tick int) money {
	policy := p.policies[inv.commodity]
	return policy.getPriceForQty(inv.qty, tick)
}

func (p *buyPolicyCatalog) registerPurchase(c commodity, price money, tick int) {
	p.policies[c].priceHistory.putPrice(price, tick)
}

type pricePolicy struct {
	priceHistory *priceHistory
	targetQty    float64
	coeff        float64
}

const windowSize = 4

func (p *pricePolicy) getPriceForQty(qty float64, tick int) money {
	pct := p.getPercentBelowTarget(qty)
	weightedPct := pct * p.coeff
	avg := p.priceHistory.average(windowSize, tick)
	if avg == 0 {
		avg = 1
	}
	price := int(avg) + int(weightedPct*float64(avg))
	return money(price)
}

func (p *pricePolicy) getPercentBelowTarget(qty float64) float64 {
	delta := p.targetQty - qty
	return delta / p.targetQty
}

func (p *pricePolicy) registerSale(price money, tick int) {
	p.priceHistory.putPrice(price, tick)
}

func (p *pricePolicy) String() string {
	return fmt.Sprintf("{priceHistory len: %d, targetQty: %g, coeff: %g}", len(p.priceHistory.prices), p.targetQty, p.coeff)
}
