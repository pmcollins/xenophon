package xenophon

type priceHistory struct {
	prices []money
}

func newPriceHistory(seed money) *priceHistory {
	h := &priceHistory{prices: nil}
	h.putPrice(seed, 0)
	return h
}

func (h *priceHistory) putPrice(price money, tick int) {
	for i := 0; i < tick-len(h.prices); i++ {
		h.prices = append(h.prices, 0)
	}
	h.prices = append(h.prices, price)
}

func (h *priceHistory) average(windowSize int, tick int) money {
	avg := money(0)
	endIdx := tick - windowSize
	if endIdx < 0 {
		endIdx = 0
	}
	for i := tick - 1; i >= endIdx; i-- {
		avg += h.priceAtIndex(i)
	}
	return avg / money(tick - endIdx)
}

func (h *priceHistory) priceAtIndex(idx int) money {
	if idx >= len(h.prices) {
		return 0
	}
	return h.prices[idx]
}
