package xenophon

import (
	"testing"
	"xenophon/internal/test"
)

func TestPriceHistory(t *testing.T) {
	h := newPriceHistory(2)
	test.AssertEqual(t, money(2), h.average(4, 1))
	h.putPrice(money(4), 1)
	test.AssertEqual(t, money(3), h.average(4, 2))
	h.putPrice(money(6), 2)
	test.AssertEqual(t, money(4), h.average(4, 3))
	h.putPrice(money(8), 3)
	test.AssertEqual(t, money(5), h.average(4, 4))
	h.putPrice(money(10), 4)
	test.AssertEqual(t, money(7), h.average(4, 5))
}

func TestNonContinuousPriceHistory(t *testing.T) {
	h := newPriceHistory(6)
	h.putPrice(money(3), 2)
	test.AssertEqual(t, money(3), h.average(3, 3))
}

func TestLateTick(t *testing.T) {
	h := newPriceHistory(6)
	test.AssertEqual(t, money(6), h.average(2, 1))
	test.AssertEqual(t, money(3), h.average(2, 2))
	test.AssertEqual(t, money(0), h.average(2, 3))
}
