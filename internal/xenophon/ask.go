package xenophon

import "fmt"

type ask struct {
	seller    *bot
	price     money
	commodity commodity
}

func (a ask) considerSaleTo(buyer *bot, tick int) money {
	if a.seller.id == buyer.id {
		return 0
	}
	bid := buyer.getBid(a.commodity, tick)
	if bid < a.price {
		return 0
	}
	salePrice := (bid + a.price) / 2
	if buyer.money < salePrice {
		return 0
	}
	a.transact(buyer, salePrice, tick)
	return salePrice
}

func (a ask) transact(buyer *bot, salePrice money, tick int) {
	fmt.Printf("[%d]<->[%d] $%v/%v\n", a.seller.id, buyer.id, salePrice, a.commodity)
	a.seller.store.contents.qty -= 1
	buyer.getPantryInv(a.commodity).qty += 1
	buyer.money -= salePrice
	a.seller.money += salePrice
	buyer.registerPurchase(a.commodity, salePrice, tick)
	a.seller.registerSale(salePrice, tick)
}
