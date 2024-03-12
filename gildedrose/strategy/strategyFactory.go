package strategy

import "kata/gildedrose/item"

func GetStrategy(item *item.Item) Strategy {
	switch item.Name {
	case "Sulfuras, Hand of Ragnaros":
		return newSulfuras(item)
	case "Aged Brie":
		return newAgedBrie(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		return newBackstagePasses(item)
	case "Conjured":
		return newConjured(item)
	default:
		return newNormal(item)
	}
}
