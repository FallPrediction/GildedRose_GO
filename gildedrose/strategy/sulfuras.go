package strategy

import "kata/gildedrose/item"

type Sulfuras struct {
	*item.Item
}

func newSulfuras(item *item.Item) *Sulfuras {
	return &Sulfuras{Item: item}
}

func (item *Sulfuras) UpdateSellIn() {
	// pass
}

func (item *Sulfuras) UpdateQuality() {
	// pass
}
