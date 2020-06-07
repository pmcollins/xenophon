package xenophon

import "fmt"

type bot struct {
	id               int
	store            *store
	pantry           *pantry
	buyPolicyCatalog *buyPolicyCatalog
	money            money
}

func newBot(id int, cmd commodity) *bot {
	return &bot{
		id:               id,
		store:            newStore(cmd),
		pantry:           newPantry(),
		buyPolicyCatalog: newBuyPolicyCatalog(),
		money:            1000,
	}
}

func (b *bot) String() string {
	return fmt.Sprintf("{id: %d, money: %v, store: %v pantry: %v}", b.id, b.money, b.store, b.pantry)
}

func (b *bot) shop(mkt *market, tick int) {
	for _, sink := range b.pantry.contents {
		ask := mkt.getLowestPriceForCommodity(sink.inv.commodity, tick)
		salePrice := ask.considerSaleTo(b, tick)
		println(salePrice)
	}
}

func (b *bot) getBid(c commodity, tick int) money {
	inv := b.getPantryInv(c)
	return b.buyPolicyCatalog.getPrice(inv, tick)
}

func (b *bot) getPantryInv(c commodity) *inventory {
	for _, sink := range b.pantry.contents {
		if sink.inv.commodity == c {
			return sink.inv
		}
	}
	return nil
}

func (b *bot) getAsk(tick int) *ask {
	ask := b.store.getAsk(tick)
	if ask == nil {
		return nil
	}
	ask.seller = b
	return ask
}

func (b *bot) consume() {
	b.pantry.consume()
}

func (b *bot) grow() {
	b.store.grow()
}

func (b *bot) registerPurchase(c commodity, price money, tick int) {
	b.buyPolicyCatalog.registerPurchase(c, price, tick)
}

func (b *bot) registerSale(price money, tick int) {
	b.store.registerSale(price, tick)
}

func (b *bot) sells(c commodity) bool {
	return b.store.contents.commodity == c
}
