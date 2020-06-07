package main

import (
	"fmt"
	"math/rand"
	"time"
	"xenophon/internal/xenophon"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	mkt := xenophon.NewMarket(10)
	for i := 1; i < 120; i++ {
		mkt.Tick(i)
		fmt.Printf("%v\n", mkt.Bot(0))
		time.Sleep(1*time.Second)
	}
}
