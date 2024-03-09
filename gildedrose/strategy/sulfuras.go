package strategy

import "kata/gildedrose/item"

type Sulfuras struct {
	*item.Item
}

func (item *Sulfuras) UpdateSellIn() {
	// pass
}

func (item *Sulfuras) UpdateQuality() {
	// pass
}
