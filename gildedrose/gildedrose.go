package gildedrose

import (
	"kata/gildedrose/item"
	"kata/gildedrose/strategy"
)

func UpdateQuality(items []*item.Item) {
	for _, item := range items {
		s := strategy.GetStrategy(item)
		s.UpdateQuality()
		s.UpdateSellIn()
	}
}
