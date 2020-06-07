package xenophon

import (
	"fmt"
	"testing"
	"xenophon/internal/test"
)

const targetQty = 40
const pantryStartQty = 10
const startMoney = 100

func TestSimplePricePolicy(t *testing.T) {
	h := newPriceHistory(100)
	p := pricePolicy{
		priceHistory: h,
		targetQty:    10,
		coeff:        1,
	}
	test.AssertEqual(t, money(100), p.getPriceForQty(10, 1))
	test.AssertEqual(t, money(90), p.getPriceForQty(11, 1))

	p = pricePolicy{
		priceHistory: newPriceHistory(100),
		targetQty:    10,
		coeff:        2,
	}
	test.AssertEqual(t, money(100), p.getPriceForQty(10, 1))
	test.AssertEqual(t, money(80), p.getPriceForQty(11, 1))
}

func TestMarket1(t *testing.T) {
	market := market{bots: []*bot{
		mkBot(1, Oats, targetQty),
		mkBot(2, Peas, targetQty),
		mkBot(3, Beans, targetQty),
		mkBot(4, Barley, targetQty),
	}}
	bot := market.bots[0]
	bid := bot.buyPolicyCatalog.getPrice(bot.pantry.contents[0].inv, 1)
	test.AssertEqual(t, money(10), bid)

	ask := bot.getAsk(1)
	test.AssertEqual(t, money(10), ask.price)
}

func TestTransact(t *testing.T) {
	buyer := mkBot(1, Oats, targetQty)
	commodity := Peas
	seller := mkBot(2, commodity, targetQty)
	ask := &ask{seller: seller, price: 10, commodity: commodity}
	salePrice := ask.considerSaleTo(buyer, 1)
	test.AssertEqual(t, float64(pantryStartQty+1), buyer.getPantryInv(commodity).qty)
	test.AssertEqual(t, startMoney-salePrice, buyer.money)
	test.AssertEqual(t, float64(targetQty-1), seller.store.contents.qty)
	test.AssertEqual(t, startMoney+salePrice, seller.money)
}

func TestLowestPrices(t *testing.T) {
	market := &market{bots: []*bot{
		mkBot(1, Oats, targetQty),
		mkBot(2, Oats, targetQty+10),
		mkBot(3, Peas, targetQty),
		mkBot(4, Peas, targetQty+10),
		mkBot(5, Beans, targetQty),
		mkBot(6, Beans, targetQty+10),
		mkBot(7, Barley, targetQty),
		mkBot(8, Barley, targetQty+10),
	}}
	prices := market.getLowestPrices(1)
	for _, ask := range prices {
		test.AssertEqual(t, money(8), ask.price)
	}
}

func TestLowestPriceForCommodity(t *testing.T) {
	market := &market{bots: []*bot{
		mkBot(1, Oats, targetQty),
		mkBot(2, Oats, targetQty+10),
		mkBot(3, Oats, targetQty-10),
		mkBot(4, Peas, targetQty+20),
	}}
	lowest := market.getLowestPriceForCommodity(Oats, 1)
	test.AssertEqual(t, 2, lowest.seller.id)
}

func TestPurchase(t *testing.T) {
	market := &market{bots: []*bot{
		mkBot(1, Oats, targetQty),
		mkBot(2, Peas, targetQty),
		mkBot(3, Beans, targetQty),
		mkBot(4, Barley, targetQty),
	}}
	buyer := market.bots[0]
	buyer.shop(market, 1)
}

func TestBotConsumeInventory(t *testing.T) {
	bot := mkBot(1, Oats, targetQty)
	fmt.Printf("%v\n", bot.pantry)
	bot.consume()
	fmt.Printf("%v\n", bot.pantry)
}

func TestGrow(t *testing.T) {
	b := &bot{
		store: &store{
			maxGrowRate: 1,
			contents:    &inventory{qty: 10},
			sellPolicy:  &pricePolicy{targetQty: 10},
		},
	}
	b.grow()
	test.AssertEqual(t, float64(10), b.store.contents.qty)
	b = &bot{
		store: &store{
			maxGrowRate: 1,
			contents:    &inventory{qty: 0},
			sellPolicy:  &pricePolicy{targetQty: 10},
		},
	}
	b.grow()
	test.AssertEqual(t, float64(1), b.store.contents.qty)
}

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\

func mkBot(id int, c commodity, qty float64) *bot {
	return &bot{
		id:               id,
		pantry:           mkPantry(),
		store:            mkStoreWithQty(c, qty),
		buyPolicyCatalog: newBuyPolicyCatalog(),
		money:            startMoney,
	}
}

func mkPantry() *pantry {
	sinks := make([]*sink, 0)
	for c := Oats; c <= Barley; c++ {
		inv := &inventory{commodity: c, qty: pantryStartQty}
		sink := &sink{inv: inv, ratePerDay: 1}
		sinks = append(sinks, sink)
	}
	return &pantry{contents: sinks}
}

func mkStoreWithQty(c commodity, qty float64) *store {
	return &store{
		contents: &inventory{commodity: c, qty: qty},
		sellPolicy: &pricePolicy{
			priceHistory: newPriceHistory(10),
			targetQty:    targetQty,
			coeff:        1,
		},
		maxGrowRate: 1,
	}
}
