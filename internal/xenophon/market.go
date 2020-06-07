package xenophon

import "math/rand"

type market struct {
	bots []*bot
}

func NewMarket(population int) *market {
	bots := make([]*bot, population)
	for i := 0; i < population; i++ {
		bots[i] = newBot(i, commodity((i%int(Barley))+1))
	}
	return &market{bots: bots}
}

func (m *market) getLowestPrices(tick int) map[commodity]*ask {
	lowest := make(map[commodity]*ask)
	for _, seller := range m.bots {
		ask := seller.getAsk(tick)
		if ask == nil {
			continue
		}
		if stored, ok := lowest[ask.commodity]; ok {
			if ask.price < stored.price {
				lowest[ask.commodity] = ask
			}
		} else {
			lowest[ask.commodity] = ask
		}
	}
	return lowest
}

func (m *market) getLowestPriceForCommodity(c commodity, tick int) *ask {
	var lowest *ask
	for _, seller := range m.bots {
		if !seller.sells(c) {
			continue
		}
		ask := seller.getAsk(tick)
		if lowest == nil || ask.price < lowest.price {
			lowest = ask
		}
	}
	return lowest
}

func (m *market) Tick(tick int) {
	if tick == 0 {
		return // illegal tick
	}
	m.consume()
	m.grow()
	m.shop(tick)
}

func (m *market) consume() {
	for _, bot := range m.bots {
		bot.consume()
	}
}

func (m *market) grow() {
	for _, bot := range m.bots {
		bot.grow()
	}
}

func (m *market) shop(tick int) {
	indexes := rand.Perm(len(m.bots))
	for _, idx := range indexes {
		bot := m.bots[idx]
		bot.shop(m, tick)
	}
}

func (m *market) Bot(i int) *bot {
	return m.bots[i]
}

func (m *market) String() string {
	var out string
	for _, bot := range m.bots {
		out += bot.String() + "\n"
	}
	return out
}
